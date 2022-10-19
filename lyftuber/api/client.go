package api

import (
	lyftapi "github.com/spudtrooper/lyft/api"
	uberapi "github.com/spudtrooper/uber/api"
)

type Client struct {
	lyftClient *lyftapi.Client
	uberClient *uberapi.Client
}

func NewClient(lyftClient *lyftapi.Client, uberClient *uberapi.Client) *Client {
	return &Client{
		lyftClient: lyftClient,
		uberClient: uberClient,
	}
}
