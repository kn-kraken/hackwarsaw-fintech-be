package db

type Business struct {
	Name     string
	Type     string
	Location Location
	Distance float32
}

type Location struct {
	Longitude float32
	Latitude  float32
}
