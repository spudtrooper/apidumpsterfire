package api

import (
	"sync"

	"github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/parallel"
	lyftapi "github.com/spudtrooper/lyft/api"
	uberapi "github.com/spudtrooper/uber/api"
)

type FaresInfo struct {
	Uber *uberapi.FareEstimateInfo `json:"uber"`
	Lyft *lyftapi.OfferingsInfo    `json:"lyft"`
}

//go:generate genopts --function Fares --params --required "lyftToken string, uberCSID string, uberSID string" originLatitude:float64:40.7701286 originLongitude:float64:-73.9829762 destinationLatitude:float64:40.7801286 destinationLongitude:float64:-73.9929762
func (c *Client) Fares(lyftToken string, uberCSID string, uberSID string, optss ...FaresOption) (*FaresInfo, error) {
	opts := MakeFaresOptions(optss...)

	uberCh := make(chan *uberapi.FareEstimateInfo)
	lyftCh := make(chan *lyftapi.OfferingsInfo)
	errsCh := make(chan error)

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()

			v, err := c.uberClient.FareEstimate(
				uberapi.FareEstimateCsid(uberCSID),
				uberapi.FareEstimateSid(uberSID),
				uberapi.FareEstimatePickupLatitude(opts.OriginLatitude()),
				uberapi.FareEstimatePickupLongitude(opts.OriginLongitude()),
				uberapi.FareEstimateDestinationLatitude(opts.DestinationLatitude()),
				uberapi.FareEstimateDestinationLongitude(opts.DestinationLongitude()),
			)
			if err != nil {
				errsCh <- err
				return
			}
			uberCh <- v
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			v, err := c.lyftClient.Offerings(
				lyftapi.OfferingsToken(lyftToken),
				lyftapi.OfferingsOriginLatitudeE6(int(opts.OriginLatitude()*1e6)),
				lyftapi.OfferingsOriginLongitudeE6(int(opts.OriginLongitude()*1e6)),
				lyftapi.OfferingsDestinationLatitudeE6(int(opts.DestinationLatitude()*1e6)),
				lyftapi.OfferingsDestinationLongitudeE6(int(opts.DestinationLongitude()*1e6)),
			)
			if err != nil {
				errsCh <- err
				return
			}
			lyftCh <- v
		}()
		wg.Wait()
		close(uberCh)
		close(lyftCh)
		close(errsCh)
	}()

	var err error
	var res FaresInfo
	parallel.WaitFor(func() {
		for d := range uberCh {
			res.Uber = d
		}
	}, func() {
		for d := range lyftCh {
			res.Lyft = d
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

	return &res, nil
}
