package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spudtrooper/minimalcli/handler"
	opentableapi "github.com/spudtrooper/opentable/api"
	opentablehandlers "github.com/spudtrooper/opentable/handlers"
	resyapi "github.com/spudtrooper/resy/api"
	resyhandlers "github.com/spudtrooper/resy/handlers"
)

func ListenAndServe(ctx context.Context, resyClient *resyapi.Client, opentableClient *opentableapi.Extended, port int, host string) error {
	var hostPort string
	if host == "localhost" {
		hostPort = fmt.Sprintf("http://localhost:%d", port)
	} else {
		hostPort = fmt.Sprintf("https://%s", host)
	}

	mux := http.NewServeMux()

	var secs []handler.Section
	{
		sec, err := handler.AddSection(ctx, mux,
			resyhandlers.CreateHandlers(resyClient),
			"resy",
			"unofficial resy API",
			handler.AddSectionFooterHTML(`<a href="/">Home</a> | Details: <a target="_" href="//github.com/spudtrooper/resy">github.com/spudtrooper/resy</a>`),
			handler.AddSectionSourceLinks(true),
			handler.AddSectionSerializedSourceLocations(resyhandlers.SourceLocations),
		)
		if err != nil {
			return err
		}
		secs = append(secs, *sec)
	}

	{
		sec, err := handler.AddSection(ctx, mux,
			opentablehandlers.CreateHandlers(opentableClient),
			"opentable",
			"unofficial opentable API",
			handler.AddSectionFooterHTML(`<a href="/">Home</a> | Details: <a target="_" href="//github.com/spudtrooper/opentable">github.com/spudtrooper/opentable</a>`),
			handler.AddSectionSourceLinks(true),
			handler.AddSectionSerializedSourceLocations(opentablehandlers.SourceLocations),
		)
		if err != nil {
			return err
		}
		secs = append(secs, *sec)
	}

	if err := handler.GenIndex(ctx, mux, secs,
		handler.GenIndexTitle("API Dumpster Fire"),
		handler.GenIndexFooterHTML(`Details: <a target="_" href="//github.com/spudtrooper/apidumpsterfire">github.com/spudtrooper/apidumpsterfire</a>`),
	); err != nil {
		return err
	}

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
