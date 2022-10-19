// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type NearbyDriversOption struct {
	f func(*nearbyDriversOptionImpl)
	s string
}

func (o NearbyDriversOption) String() string { return o.s }

type NearbyDriversOptions interface {
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
}

func NearbyDriversLatitude(latitude float64) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("api.NearbyDriversLatitude(float64 %+v)}", latitude)}
}
func NearbyDriversLatitudeFlag(latitude *float64) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("api.NearbyDriversLatitude(float64 %+v)}", latitude)}
}

func NearbyDriversLongitude(longitude float64) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("api.NearbyDriversLongitude(float64 %+v)}", longitude)}
}
func NearbyDriversLongitudeFlag(longitude *float64) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("api.NearbyDriversLongitude(float64 %+v)}", longitude)}
}

type nearbyDriversOptionImpl struct {
	latitude      float64
	has_latitude  bool
	longitude     float64
	has_longitude bool
}

func (n *nearbyDriversOptionImpl) Latitude() float64  { return n.latitude }
func (n *nearbyDriversOptionImpl) HasLatitude() bool  { return n.has_latitude }
func (n *nearbyDriversOptionImpl) Longitude() float64 { return n.longitude }
func (n *nearbyDriversOptionImpl) HasLongitude() bool { return n.has_longitude }

type NearbyDriversParams struct {
	LyftToken string  `json:"lyft_token" required:"true"`
	UberCSID  string  `json:"uber_csid" required:"true"`
	UberSID   string  `json:"uber_sid" required:"true"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (o NearbyDriversParams) Options() []NearbyDriversOption {
	return []NearbyDriversOption{
		NearbyDriversLatitude(o.Latitude),
		NearbyDriversLongitude(o.Longitude),
	}
}

func makeNearbyDriversOptionImpl(opts ...NearbyDriversOption) *nearbyDriversOptionImpl {
	res := &nearbyDriversOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeNearbyDriversOptions(opts ...NearbyDriversOption) NearbyDriversOptions {
	return makeNearbyDriversOptionImpl(opts...)
}
