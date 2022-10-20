package api

import (
	"sync"

	"github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/parallel"
	lyftapi "github.com/spudtrooper/lyft/api"
	uberapi "github.com/spudtrooper/uber/api"
)

type NearbyDriversInfoDriver struct {
	ID        string                    `json:"id"`
	Type      string                    `json:"type"`
	Latitude  float64                   `json:"latitude"`
	Longitude float64                   `json:"longitude"`
	ImageURL  string                    `json:"image_url"`
	RootInfo  NearbyDriversInfoRootInfo `json:"info"`
}

type NearbyDriversInfoRootInfo struct {
	Uber *uberapi.StatusInfo        `json:"uber"`
	Lyft *lyftapi.NearbyDriversInfo `json:"lyft"`
}

type NearbyDriversInfo struct {
	Drivers []NearbyDriversInfoDriver
}

//go:generate genopts --function NearbyDrivers --params --required "lyftToken string, uberCSID string, uberSID string" latitude:float64:40.7701286 longitude:float64:-73.9829762
func (c *Client) NearbyDrivers(lyftToken string, uberCSID string, uberSID string, optss ...NearbyDriversOption) (*NearbyDriversInfo, error) {
	opts := MakeNearbyDriversOptions(optss...)

	// var drivers []NearbyDriversInfoDriver

	driversCh := make(chan NearbyDriversInfoDriver)
	errsCh := make(chan error)

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()

			v, err := c.uberClient.Status(
				uberapi.StatusCsid(uberCSID),
				uberapi.StatusSid(uberSID),
				uberapi.StatusLatitude(opts.Latitude()),
				uberapi.StatusLongitude(opts.Longitude()),
			)
			if err != nil {
				errsCh <- err
				return
			}
			vehicleViews, err := v.VehicleViews()
			if err != nil {
				errsCh <- err
				return
			}
			for _, vv := range vehicleViews {
				d := NearbyDriversInfoDriver{
					Type:      "uber",
					ID:        vv.ID,
					Latitude:  vv.Lat,
					Longitude: vv.Lng,
					ImageURL:  vv.ImageURL,
					RootInfo: NearbyDriversInfoRootInfo{
						Uber: v,
					},
				}
				driversCh <- d
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			v, err := c.lyftClient.NearbyDrivers(
				lyftapi.NearbyDriversToken(lyftToken),
				lyftapi.NearbyDriversDestinationLatitudeE6(int(opts.Latitude()*1e6)),
				lyftapi.NearbyDriversDestinationLongitudeE6(int(opts.Longitude()*1e6)),
			)
			if err != nil {
				errsCh <- err
				return
			}
			vehicleViews, err := v.VehicleViews()
			if err != nil {
				errsCh <- err
				return
			}
			for _, vv := range vehicleViews {
				d := NearbyDriversInfoDriver{
					Type:      "lyft",
					ID:        vv.ID,
					Latitude:  vv.Lat,
					Longitude: vv.Lng,
					ImageURL:  vv.ImageURL,
					RootInfo: NearbyDriversInfoRootInfo{
						Lyft: v,
					},
				}
				driversCh <- d
			}
		}()
		wg.Wait()
		close(driversCh)
		close(errsCh)
	}()

	var drivers []NearbyDriversInfoDriver
	var err error
	parallel.WaitFor(func() {
		for d := range driversCh {
			drivers = append(drivers, d)
		}
	}, func() {
		b := errors.MakeErrorCollector()
		for e := range errsCh {
			b.Add(e)
		}
		err = b.Build()
	})

	if err != nil {
		return nil, err
	}

	res := &NearbyDriversInfo{
		Drivers: drivers,
	}
	return res, nil
}
