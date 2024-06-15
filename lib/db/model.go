package db

type Business struct {
	Name     string
	Type     BusinessType `binding:"enum"`
	Location Location
	Distance float32
}

type Location struct {
	Longitude float32
	Latitude  float32
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
