package api

import (
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

	var drivers []NearbyDriversInfoDriver

	{
		v, err := c.uberClient.Status(
			uberapi.StatusCsid(uberCSID),
			uberapi.StatusSid(uberSID),
			uberapi.StatusLatitude(opts.Latitude()),
			uberapi.StatusLongitude(opts.Longitude()),
		)
		if err != nil {
			return nil, err
		}
		vehicleViews, err := v.VehicleViews()
		if err != nil {
			return nil, err
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
			drivers = append(drivers, d)
		}
	}

	{
		v, err := c.lyftClient.NearbyDrivers(
			lyftapi.NearbyDriversToken(lyftToken),
			lyftapi.NearbyDriversDestinationLatitudeE6(int(opts.Latitude()/100.0)),
			lyftapi.NearbyDriversDestinationLongitudeE6(int(opts.Longitude()/100.0)),
		)
		if err != nil {
			return nil, err
		}
		vehicleViews, err := v.VehicleViews()
		if err != nil {
			return nil, err
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
			drivers = append(drivers, d)
		}
	}

	res := &NearbyDriversInfo{
		Drivers: drivers,
	}
	return res, nil
}
