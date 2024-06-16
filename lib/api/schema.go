package api

import (
	"slices"

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
	Distance float32         `json:"distance"`
}

type Location struct {
	Longitude float32 `json:"lng"`
	Latitude  float32 `json:"lat"`
}

type RealEstate struct {
	Address  string   `json:"address"`
	Score    float32  `json:"score"`
	Area     float32  `json:"area"`
	Location Location `json:"location"`
	Distance float32  `json:"distance"`
}

func NewRealEstateScores(businesses []db.Business, realEstates []db.RealEstate, scores []float32) RealEstateScores {
	businessesDto := utils.MapRef(businesses, NewBusiness)
	realEstatesDto := utils.MapRef2(realEstates, scores, NewRealEstate)
  slices.SortFunc(realEstatesDto, func(first RealEstate, second RealEstate) int {
    if (first.Score > second.Score) {
      return -1
    } else if (first.Score == second.Score) {
      return 0
    } else {
      return 1
    }
  })
	return RealEstateScores{
		Businesses:  businessesDto,
		RealEstates: realEstatesDto,
	}
}

func NewBusiness(business *db.Business) Business {
	return Business{
		Name:     business.Name,
		Type:     business.Type,
		Location: NewLocation(business.Location),
		Address:  "ul. Długa 123",
		Distance: business.Distance,
	}
}

func NewLocation(location db.Location) Location {
	return Location{
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	}
}

func NewRealEstate(realEstate *db.RealEstate, score *float32) RealEstate {
	return RealEstate{
		Address:  realEstate.Address,
		Score:    *score,
		Area:     realEstate.Area,
		Location: NewLocation(realEstate.Location),
		Distance: realEstate.Distance,
	}
}
