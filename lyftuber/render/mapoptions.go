// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package render

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type MapOption struct {
	f func(*mapOptionImpl)
	s string
}

func (o MapOption) String() string { return o.s }

type MapOptions interface {
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	SleepSecs() int
	HasSleepSecs() bool
	Zoom() int
	HasZoom() bool
}

func MapLatitude(latitude float64) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("render.MapLatitude(float64 %+v)", latitude)}
}
func MapLatitudeFlag(latitude *float64) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("render.MapLatitude(float64 %+v)", latitude)}
}

func MapLongitude(longitude float64) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("render.MapLongitude(float64 %+v)", longitude)}
}
func MapLongitudeFlag(longitude *float64) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("render.MapLongitude(float64 %+v)", longitude)}
}

func MapSleepSecs(sleepSecs int) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		opts.has_sleepSecs = true
		opts.sleepSecs = sleepSecs
	}, fmt.Sprintf("render.MapSleepSecs(int %+v)", sleepSecs)}
}
func MapSleepSecsFlag(sleepSecs *int) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		if sleepSecs == nil {
			return
		}
		opts.has_sleepSecs = true
		opts.sleepSecs = *sleepSecs
	}, fmt.Sprintf("render.MapSleepSecs(int %+v)", sleepSecs)}
}

func MapZoom(zoom int) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		opts.has_zoom = true
		opts.zoom = zoom
	}, fmt.Sprintf("render.MapZoom(int %+v)", zoom)}
}
func MapZoomFlag(zoom *int) MapOption {
	return MapOption{func(opts *mapOptionImpl) {
		if zoom == nil {
			return
		}
		opts.has_zoom = true
		opts.zoom = *zoom
	}, fmt.Sprintf("render.MapZoom(int %+v)", zoom)}
}

type mapOptionImpl struct {
	latitude      float64
	has_latitude  bool
	longitude     float64
	has_longitude bool
	sleepSecs     int
	has_sleepSecs bool
	zoom          int
	has_zoom      bool
}

func (m *mapOptionImpl) Latitude() float64  { return or.Float64(m.latitude, 40.7701286) }
func (m *mapOptionImpl) HasLatitude() bool  { return m.has_latitude }
func (m *mapOptionImpl) Longitude() float64 { return or.Float64(m.longitude, -73.9829762) }
func (m *mapOptionImpl) HasLongitude() bool { return m.has_longitude }
func (m *mapOptionImpl) SleepSecs() int     { return or.Int(m.sleepSecs, 0) }
func (m *mapOptionImpl) HasSleepSecs() bool { return m.has_sleepSecs }
func (m *mapOptionImpl) Zoom() int          { return or.Int(m.zoom, 14) }
func (m *mapOptionImpl) HasZoom() bool      { return m.has_zoom }

type MapParams struct {
	Latitude  float64 `json:"latitude" default:"40.7701286"`
	Longitude float64 `json:"longitude" default:"-73.9829762"`
	LyftToken string  `json:"lyft_token" required:"true"`
	SleepSecs int     `json:"sleep_secs" default:"0"`
	UberCSID  string  `json:"uber_csid" required:"true"`
	UberSID   string  `json:"uber_sid" required:"true"`
	Zoom      int     `json:"zoom" default:"14"`
}

func (o MapParams) Options() []MapOption {
	return []MapOption{
		MapLatitude(o.Latitude),
		MapLongitude(o.Longitude),
		MapSleepSecs(o.SleepSecs),
		MapZoom(o.Zoom),
	}
}

func makeMapOptionImpl(opts ...MapOption) *mapOptionImpl {
	res := &mapOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeMapOptions(opts ...MapOption) MapOptions {
	return makeMapOptionImpl(opts...)
}
