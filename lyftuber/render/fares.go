package render

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/spudtrooper/apidumpsterfire/lyftuber/api"
	lyftapi "github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
	uberapi "github.com/spudtrooper/uber/api"
)

//go:embed tmpl/fares.html
var faresTmpl string

func Fares(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.FaresInfo)

	config := handler.RendererConfig{
		IsFragment: true,
	}

	type row struct {
		ID               string
		Type             string
		Description      string
		Fare             string
		EstimatedMinutes int
		Multiplier       string
		JSON             interface{}
	}
	var uberRows, lyftRows []row

	for typ, f := range params.Uber.FareEstimate.Fares {
		if typ == "" || f.FareUUID == "" {
			continue
		}
		var jsonObj = struct {
			Fare uberapi.FareEstimateInfoFareEstimateFare
		}{
			Fare: f,
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, config, err
		}
		jsonStr := string(jsonBytes)

		uberRows = append(uberRows, row{
			ID:               f.ProductUUID,
			Type:             typ,
			Description:      "--",
			Fare:             fmt.Sprintf("%s %s", f.Fare, f.CurrencyCode),
			EstimatedMinutes: f.EstimatedTripTime,
			Multiplier:       fmt.Sprintf("%0.2f", f.UpfrontFare.SurgeMultiplier),
			JSON:             jsonStr,
		})
	}

	for _, o := range params.Lyft.Offers {
		var jsonObj = struct {
			Fare lyftapi.OfferingsInfoOffer
		}{
			Fare: o,
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, config, err
		}
		jsonStr := string(jsonBytes)

		var fare, desc string
		var estMins int
		if len(o.CostEstimate.LineItems) > 0 {
			amount, err := strconv.Atoi(o.CostEstimate.LineItems[0].Amount)
			if err != nil {
				return nil, config, err
			}
			fare = fmt.Sprintf("%0.2f %s", float64(amount)/100.0, o.CostEstimate.LineItems[0].Currency)
			desc = o.CostEstimate.LineItems[0].Description
			var secs int
			if o.CostEstimate.EstimatedDurationSeconds != "" {
				s, err := strconv.Atoi(o.CostEstimate.EstimatedDurationSeconds)
				if err != nil {
					return nil, config, err
				}
				secs = s
			}
			estMins = secs / 60
		}
		lyftRows = append(lyftRows, row{
			ID:               o.ID,
			Type:             o.OfferProductID,
			Description:      desc,
			Fare:             fare,
			EstimatedMinutes: estMins,
			Multiplier:       fmt.Sprintf("%0.2f", o.CostEstimate.PrimetimeMultiplier),
			JSON:             jsonStr,
		})
	}

	sort.Slice(uberRows, func(i, j int) bool { return uberRows[i].ID < uberRows[j].ID })
	sort.Slice(lyftRows, func(i, j int) bool { return lyftRows[i].ID < lyftRows[j].ID })
	var data = struct {
		Uber []row
		Lyft []row
	}{
		Uber: uberRows,
		Lyft: lyftRows,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, faresTmpl, "Fares", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
