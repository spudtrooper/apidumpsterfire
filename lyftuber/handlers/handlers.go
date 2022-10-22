// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/apidumpsterfire/lyftuber/api"
	"github.com/spudtrooper/apidumpsterfire/lyftuber/render"
	"github.com/spudtrooper/minimalcli/handler"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/apidumpsterfire/blob/main/lyftuber/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Client) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("NearbyDrivers",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.NearbyDriversParams)
			return client.NearbyDrivers(p.LyftToken, p.UberCSID, p.UberSID, p.Options()...)
		},
		api.NearbyDriversParams{},
		handler.NewHandlerRenderer(render.NearbyDrivers),
	)

	b.NewStaticHandler("Map",
		render.MapTmpl,
		render.MapParams{},
		handler.NewHandlerRendererConfig(handler.RendererConfig{IsFragment: false}),
	)

	b.NewHandler("Fares",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.FaresParams)
			return client.Fares(p.LyftToken, p.UberCSID, p.UberSID, p.Options()...)
		},
		api.FaresParams{},
		handler.NewHandlerRenderer(render.Fares),
	)

	return b.Build()
}
