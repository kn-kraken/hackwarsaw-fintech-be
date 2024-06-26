package db

type RealEstate = struct {
	Id           string   `json:"id"`
	Address      string   `json:"address"`
	OccuanceType string   `json:"occurance_type"`
	Area         float32  `json:"area"`
	InitialPrice float32  `json:"initial_price"`
	District     string   `json:"district"`
	Location     Location `json:"location"`
	Distance     float32  `json:"distance"`
}

type Polygon struct {
	Id        int        `json:"osm_int"`
	Name      string     `json:"name"`
	Locations []Location `json:"coordinates"`
}

type Business struct {
	Name     string
	Type     BusinessType `binding:"enum"`
	Location Location
	Distance float32
}

type Location struct {
	Longitude float32 `json:"lng"`
	Latitude  float32 `json:"lat"`
}

type BusinessType string

const (
	ALCOHOL      BusinessType = "ALCOHOL"
	BAKERY       BusinessType = "BAKERY"
	BAR          BusinessType = "BAR"
	BUTCHER      BusinessType = "BUTCHER"
	CAFE         BusinessType = "CAFE"
	ELECTRONICS  BusinessType = "ELECTRONICS"
	GREENGROCER  BusinessType = "GREENGROCER"
	HAIRDRESSER  BusinessType = "HAIRDRESSER"
	LOCKSMITH    BusinessType = "LOCKSMITH"
	PET_GROOMING BusinessType = "PET_GROOMING"
	RESTAURANT   BusinessType = "RESTAURANT"
	SHOE_REPAIR  BusinessType = "SHOE_REPAIR"
	TAILOR       BusinessType = "TAILOR"
)

func (r BusinessType) IsValid() bool {
	switch r {
	case

		ALCOHOL,
		BAKERY,
		BAR,
		BUTCHER,
		CAFE,
		ELECTRONICS,
		GREENGROCER,
		HAIRDRESSER,
		LOCKSMITH,
		PET_GROOMING,
		RESTAURANT,
		SHOE_REPAIR,
		TAILOR:
		return true
	}

	return false
}
