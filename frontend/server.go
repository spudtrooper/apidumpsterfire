package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	lyftapi "github.com/spudtrooper/lyft/api"
	lyfthandlers "github.com/spudtrooper/lyft/handlers"
	"github.com/spudtrooper/minimalcli/handler"
	opensecretsapi "github.com/spudtrooper/opensecrets/api"
	opensecretshandlers "github.com/spudtrooper/opensecrets/handlers"
	opentableapi "github.com/spudtrooper/opentable/api"
	opentablehandlers "github.com/spudtrooper/opentable/handlers"
	resyapi "github.com/spudtrooper/resy/api"
	resyhandlers "github.com/spudtrooper/resy/handlers"
)

func ListenAndServe(ctx context.Context,
	resyClient *resyapi.Client,
	opentableClient *opentableapi.Extended,
	opensecretsClient *opensecretsapi.Core,
	lyftClient *lyftapi.Client,
	port int, host string) error {
	var hostPort string
	if host == "localhost" {
		hostPort = fmt.Sprintf("http://localhost:%d", port)
	} else {
		hostPort = fmt.Sprintf("https://%s", host)
	}

	mux := http.NewServeMux()
	handler.Init(mux)

	var secs []handler.Section

	{
		sec, err := handler.AddSection(ctx, mux,
			opentablehandlers.CreateHandlers(opentableClient),
			"opentable",
			"unofficial opentable API",
			handler.AddSectionKey("opentable"),
			handler.AddSectionFooterHTML(`<a href="/">Home</a> | Details: <a target="_" href="//github.com/spudtrooper/opentable">github.com/spudtrooper/opentable</a>`),
			handler.AddSectionSourceLinks(true),
			handler.AddSectionSerializedSourceLocations(opentablehandlers.SourceLocations),
		)
		if err != nil {
			return err
		}
		secs = append(secs, *sec)
	}

	{
		sec, err := handler.AddSection(ctx, mux,
			opensecretshandlers.CreateHandlers(opensecretsClient),
			"opensecrets",
			"unofficial opensecrets API",
			handler.AddSectionKey("opensecrets"),
			handler.AddSectionFooterHTML(`<a href="/">Home</a> | Details: <a target="_" href="//github.com/spudtrooper/opensecrets">github.com/spudtrooper/opensecrets</a>`),
			handler.AddSectionSourceLinks(true),
			handler.AddSectionSerializedSourceLocations(opensecretshandlers.SourceLocations),
		)
		if err != nil {
			return err
		}
		secs = append(secs, *sec)
	}

	{
		sec, err := handler.AddSection(ctx, mux,
			lyfthandlers.CreateHandlers(lyftClient),
			"lyft",
			"unofficial lyft API",
			handler.AddSectionKey("lyft"),
			handler.AddSectionFooterHTML(`<a href="/">Home</a> | Details: <a target="_" href="//github.com/spudtrooper/lyft">github.com/spudtrooper/lyft</a>`),
			handler.AddSectionSourceLinks(true),
			handler.AddSectionSerializedSourceLocations(lyfthandlers.SourceLocations),
		)
		if err != nil {
			return err
		}
		secs = append(secs, *sec)
	}

	{
		sec, err := handler.AddSection(ctx, mux,
			resyhandlers.CreateHandlers(resyClient),
			"resy",
			"unofficial resy API",
			handler.AddSectionKey("resy"),
			handler.AddSectionFooterHTML(`<a href="/">Home</a> | Details: <a target="_" href="//github.com/spudtrooper/resy">github.com/spudtrooper/resy</a>`),
			handler.AddSectionSourceLinks(true),
			handler.AddSectionSerializedSourceLocations(resyhandlers.SourceLocations),
		)
		if err != nil {
			return err
		}
		secs = append(secs, *sec)
	}

	if err := handler.GenIndex(ctx, mux, secs,
		handler.GenIndexTitle("API Dumpster Fire"),
		handler.GenIndexFooterHTML(`<a href="/_all">All</a> | Details: <a target="_" href="//github.com/spudtrooper/apidumpsterfire">github.com/spudtrooper/apidumpsterfire</a>`),
	); err != nil {
		return err
	}

	if err := handler.GenAll(ctx, mux, secs,
		handler.GenAllTitle("API Dumpster Fire"),
		handler.GenAllFooterHTML(`<a href="/">Home</a> | Details: <a target="_" href="//github.com/spudtrooper/apidumpsterfire">github.com/spudtrooper/apidumpsterfire</a>`),
	); err != nil {
		return err
	}

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
