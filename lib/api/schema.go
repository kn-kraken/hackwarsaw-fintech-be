package api

import "github.com/kn-kraken/hackwarsaw-fintech/lib/db"

type Business struct {
	Name     string
	Type     string
	Location Location
	Address  string
}

type Location struct {
	Longitude float32
	Latitude  float32
}

func NewBusiness(business *db.Business) Business {
	return Business{
		Name:     business.Name,
		Type:     business.Type,
		Location: NewLocation(business.Location),
		Address:  "ul. Długa 123",
	}
}

func NewLocation(location db.Location) Location {
	return Location{
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}
}
