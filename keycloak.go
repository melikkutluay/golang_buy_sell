package main

import (
	"github.com/Nerzal/gocloak/v7"
)

type keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

func newKeycloak() *keycloak {

	return &keycloak{
		gocloak:      gocloak.NewClient("http://localhost:8088/keycloak/"),
		clientId:     "my-go-service",
		clientSecret: "dafe6352-1e63-48dc-9f87-8320e650177f",
		realm:        "medium",
	}
}
