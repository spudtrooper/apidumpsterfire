// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type FaresOption struct {
	f func(*faresOptionImpl)
	s string
}

func (o FaresOption) String() string { return o.s }

type FaresOptions interface {
	DestinationLatitude() float64
	HasDestinationLatitude() bool
	DestinationLongitude() float64
	HasDestinationLongitude() bool
	OriginLatitude() float64
	HasOriginLatitude() bool
	OriginLongitude() float64
	HasOriginLongitude() bool
}

func FaresDestinationLatitude(destinationLatitude float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		opts.has_destinationLatitude = true
		opts.destinationLatitude = destinationLatitude
	}, fmt.Sprintf("api.FaresDestinationLatitude(float64 %+v)", destinationLatitude)}
}
func FaresDestinationLatitudeFlag(destinationLatitude *float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		if destinationLatitude == nil {
			return
		}
		opts.has_destinationLatitude = true
		opts.destinationLatitude = *destinationLatitude
	}, fmt.Sprintf("api.FaresDestinationLatitude(float64 %+v)", destinationLatitude)}
}

func FaresDestinationLongitude(destinationLongitude float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		opts.has_destinationLongitude = true
		opts.destinationLongitude = destinationLongitude
	}, fmt.Sprintf("api.FaresDestinationLongitude(float64 %+v)", destinationLongitude)}
}
func FaresDestinationLongitudeFlag(destinationLongitude *float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		if destinationLongitude == nil {
			return
		}
		opts.has_destinationLongitude = true
		opts.destinationLongitude = *destinationLongitude
	}, fmt.Sprintf("api.FaresDestinationLongitude(float64 %+v)", destinationLongitude)}
}

func FaresOriginLatitude(originLatitude float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		opts.has_originLatitude = true
		opts.originLatitude = originLatitude
	}, fmt.Sprintf("api.FaresOriginLatitude(float64 %+v)", originLatitude)}
}
func FaresOriginLatitudeFlag(originLatitude *float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		if originLatitude == nil {
			return
		}
		opts.has_originLatitude = true
		opts.originLatitude = *originLatitude
	}, fmt.Sprintf("api.FaresOriginLatitude(float64 %+v)", originLatitude)}
}

func FaresOriginLongitude(originLongitude float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		opts.has_originLongitude = true
		opts.originLongitude = originLongitude
	}, fmt.Sprintf("api.FaresOriginLongitude(float64 %+v)", originLongitude)}
}
func FaresOriginLongitudeFlag(originLongitude *float64) FaresOption {
	return FaresOption{func(opts *faresOptionImpl) {
		if originLongitude == nil {
			return
		}
		opts.has_originLongitude = true
		opts.originLongitude = *originLongitude
	}, fmt.Sprintf("api.FaresOriginLongitude(float64 %+v)", originLongitude)}
}

type faresOptionImpl struct {
	destinationLatitude      float64
	has_destinationLatitude  bool
	destinationLongitude     float64
	has_destinationLongitude bool
	originLatitude           float64
	has_originLatitude       bool
	originLongitude          float64
	has_originLongitude      bool
}

func (f *faresOptionImpl) DestinationLatitude() float64 {
	return or.Float64(f.destinationLatitude, 40.7801286)
}
func (f *faresOptionImpl) HasDestinationLatitude() bool { return f.has_destinationLatitude }
func (f *faresOptionImpl) DestinationLongitude() float64 {
	return or.Float64(f.destinationLongitude, -73.9929762)
}
func (f *faresOptionImpl) HasDestinationLongitude() bool { return f.has_destinationLongitude }
func (f *faresOptionImpl) OriginLatitude() float64       { return or.Float64(f.originLatitude, 40.7701286) }
func (f *faresOptionImpl) HasOriginLatitude() bool       { return f.has_originLatitude }
func (f *faresOptionImpl) OriginLongitude() float64 {
	return or.Float64(f.originLongitude, -73.9829762)
}
func (f *faresOptionImpl) HasOriginLongitude() bool { return f.has_originLongitude }

type FaresParams struct {
	DestinationLatitude  float64 `json:"destination_latitude" default:"40.7801286"`
	DestinationLongitude float64 `json:"destination_longitude" default:"-73.9929762"`
	LyftToken            string  `json:"lyft_token" required:"true"`
	OriginLatitude       float64 `json:"origin_latitude" default:"40.7701286"`
	OriginLongitude      float64 `json:"origin_longitude" default:"-73.9829762"`
	UberCSID             string  `json:"uber_csid" required:"true"`
	UberSID              string  `json:"uber_sid" required:"true"`
}

func (o FaresParams) Options() []FaresOption {
	return []FaresOption{
		FaresDestinationLatitude(o.DestinationLatitude),
		FaresDestinationLongitude(o.DestinationLongitude),
		FaresOriginLatitude(o.OriginLatitude),
		FaresOriginLongitude(o.OriginLongitude),
	}
}

func makeFaresOptionImpl(opts ...FaresOption) *faresOptionImpl {
	res := &faresOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeFaresOptions(opts ...FaresOption) FaresOptions {
	return makeFaresOptionImpl(opts...)
}
