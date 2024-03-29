package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/spudtrooper/apidumpsterfire/frontend"
	"github.com/spudtrooper/goutil/check"

	lyftuberapi "github.com/spudtrooper/apidumpsterfire/lyftuber/api"
	lyftapi "github.com/spudtrooper/lyft/api"
	opensecretsapi "github.com/spudtrooper/opensecrets/api"
	opentableapi "github.com/spudtrooper/opentable/api"
	resyapi "github.com/spudtrooper/resy/api"
	spotifydownapi "github.com/spudtrooper/spotifydown/api"
	uberapi "github.com/spudtrooper/uber/api"
)

var (
	portForTesting = flag.Int("port_for_testing", 0, "port to listen on")
	host           = flag.String("host", "api-dumpdster-fire.herokuapp.com", "host name")
)

func main() {
	flag.Parse()
	var port int
	if *portForTesting != 0 {
		port = *portForTesting
	} else {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("invalid port: %v", err)
		}
		port = p
	}
	if port == 0 {
		log.Fatalf("port is required")
	}
	if *host == "" {
		log.Fatalf("host is required")
	}
	ctx := context.Background()

	resy := resyapi.NewClient("")
	opentable := opentableapi.FromClient(opentableapi.NewClient(""), opentableapi.EmptyCache())
	opensecrets := opensecretsapi.NewClient("")
	lyft := lyftapi.NewClient("")
	uber := uberapi.NewClient("", "")
	lyftuber := lyftuberapi.NewClient(lyft, uber)
	spotifydown := spotifydownapi.NewClient()

	check.Err(frontend.ListenAndServe(ctx,
		resy, opentable, opensecrets, lyft, uber, lyftuber, spotifydown,
		port, *host))
}
