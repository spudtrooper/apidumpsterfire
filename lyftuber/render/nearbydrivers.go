package render

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"sort"

	"github.com/spudtrooper/apidumpsterfire/lyftuber/api"
	"github.com/spudtrooper/minimalcli/handler"
)

//go:embed tmpl/nearby-drivers.html
var nearbyDriversTmpl string

func NearbyDrivers(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.NearbyDriversInfo)

	config := handler.RendererConfig{
		IsFragment: true,
	}

	type vehicleView struct {
		Type     string
		ID       string
		ImageURL string
		Lat, Lng float64
		JSON     string
	}
	var vehicleViews []vehicleView

	for _, d := range params.Drivers {
		var jsonObj = struct {
			Info api.NearbyDriversInfoDriver
		}{
			Info: d,
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, config, err
		}
		jsonStr := string(jsonBytes)

		vehicleViews = append(vehicleViews, vehicleView{
			Type:     d.Type,
			ID:       d.ID,
			ImageURL: d.ImageURL,
			Lat:      d.Latitude,
			Lng:      d.Longitude,
			JSON:     jsonStr,
		})
	}

	sort.Slice(vehicleViews, func(i, j int) bool { return vehicleViews[i].ID < vehicleViews[j].ID })
	var data = struct {
		VehicleViews []vehicleView
	}{
		VehicleViews: vehicleViews,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, nearbyDriversTmpl, "NearbyDrivers", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
