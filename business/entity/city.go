package entity

import (
	"strconv"

	"gitlab.com/hookactions/gqlgen-relay/relay"
)

func (CityConnection) IsConnection() {}

func (CityEdge) IsEdge() {}

//go:generate sh -c "go run gitlab.com/hookactions/gqlgen-relay -name City -pkg entity -type *City -cursor > city_relay.go"
type City struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

func (City) IsNode() {
}

func (c City) GetID() relay.ID {
	return c
}

func (c City) String() string {
	return strconv.FormatInt(int64(c.ID), 10)
}

type ParamCityInput struct {
	ID []*int `json:"id"`
}
