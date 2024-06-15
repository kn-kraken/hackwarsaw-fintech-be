package api

import (
	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/utils"
)

type RealEstateScoresRequest struct {
	BusinessType db.BusinessType `form:"business" binding:"required,enum"`
	Longitude    float32         `form:"longitude" binding:"required,gte=-180,lte=180"`
	Latitude     float32         `form:"latitude" binding:"required,gte=-90,lte=90"`
	Distance     float32         `form:"range" binding:"required,gt=0"`
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
	Score    float32  `json:"score"`
	Area     float32  `json:"area"`
	Location Location `json:"location"`
}

func NewRealEstateScores(businesses []db.Business) RealEstateScores {
	businessesDto := utils.MapRef(businesses, NewBusiness)
	return RealEstateScores{
		Businesses: businessesDto,
		RealEstates: []RealEstate{
			{
				Address: "Dupa",
				Score:   0.53,
				Area:    93.12,
				Location: Location{
					Longitude: 21.13893,
					Latitude:  52.280427,
				},
			},
			{
				Address: "Dupa",
				Score:   0.21,
				Area:    54.22,
				Location: Location{
					Longitude: 21.13993,
					Latitude:  52.280327,
				},
			},
		},
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
