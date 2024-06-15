package api

import (
	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/utils"
)

type RealEstateScoresRequest struct {
  BusinessType db.BusinessType `form:"business" binding:"required,enum"`
}

type RealEstateScores struct {
	Businesses  []Business   `json:"businesses"`
	RealEstates []RealEstate `json:"real_estates"`
}

type Business struct {
	Name     string          `json:"name"`
	Type     db.BusinessType `json:"type"`
	Location Location        `json:"location"`
	Address  string          `json:"address"`
}

type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type RealEstate struct {
	Address  string   `json:"address"`
	Score    Location `json:"score"`
	Area     float32  `json:"area"`
	Location Location `json:"location"`
}

func NewRealEstateScores(businesses []db.Business) RealEstateScores {
	businessesDto := utils.MapRef(businesses, NewBusiness)
	return RealEstateScores{
		Businesses:  businessesDto,
		RealEstates: []RealEstate{},
	}
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
