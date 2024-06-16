package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin/binding"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/mapaum"
	_ "github.com/lib/pq"
	"golang.org/x/exp/maps"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "gis"
	user     = "gisuser"
	password = "gispassword"
)

type Database struct {
	db *sql.DB
}

func New() (*Database, error) {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, fmt.Errorf("opening db connection: %w", err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("pinging db: %w", err)
	}

	slog.Info("connected to db")

	database := &Database{db: db}

	return database, nil
}

func (r *Database) Close() {
	r.db.Close()
}

func (r *Database) ListBusinessesInArea(
	businessType BusinessType,
	longitude float32,
	latitude float32,
	distance float32,
) ([]Business, error) {
	const query = `
SELECT
    name,
    UPPER(COALESCE(shop, amenity)) AS type,
    ST_X(ST_Transform(way, 4326)) AS lon,
    ST_Y(ST_Transform(way, 4326)) AS lat,
    ST_Distance(
        ST_SetSRID(ST_MakePoint($1, $2), 4326),
        ST_Transform(way, 4326)
    ) AS distance
FROM planet_osm_point
WHERE
    ST_DWithin(
        ST_SetSRID(ST_MakePoint($1, $2), 4326),
        ST_Transform(way, 4326),
        $3
    )
    AND name != ''
    AND UPPER(shop) = $4
    -- AND UPPER(shop) IN (
    --   'ALCOHOL',
    --   'BAKERY',
    --   'BAR',
    --   'BUTCHER',
    --   'CAFE',
    --   'ELECTRONICS',
    --   'GREENGROCER',
    --   'HAIRDRESSER',
    --   'LOCKSMITH',
    --   'PET_GROOMING',
    --   'RESTAURANT',
    --   'SHOE_REPAIR',
    --   'TAILOR'
    -- )
ORDER BY distance;
`

	rows, err := r.db.Query(query,
		fmt.Sprint(longitude),
		fmt.Sprint(latitude),
		fmt.Sprint(distance*0.01),
		businessType,
	)
	if err != nil {
		return nil, fmt.Errorf("running query: %w", err)
	}
	defer rows.Close()

	var result []Business
	for rows.Next() {
		var business Business

		err = rows.Scan(&business.Name,
			&business.Type,
			&business.Location.Longitude,
			&business.Location.Latitude,
			&business.Distance,
		)
		if err != nil {
			slog.Error("scanning row", err)
		}

		err = binding.Validator.ValidateStruct(business)
		if err != nil {
			slog.Error("validating", "error", err)
			continue
		}

		result = append(result, business)
	}

	return result, nil
}

func (r *Database) AddRealEstate(realEstate *mapaum.RealEstate) error {
	const statement = `
INSERT INTO real_estates
  (address, occurance_type, area, initial_price, district, longitude, latitude)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (address)
  DO UPDATE SET longitude=EXCLUDED.longitude, latitude=EXCLUDED.latitude
RETURNING real_estate_id
`
	var id int
	err := r.db.QueryRow(
		statement,
		realEstate.Address,
		realEstate.OccuanceType,
		realEstate.Area,
		realEstate.InitialPrice,
		realEstate.District,
		realEstate.Longitude,
		realEstate.Latitude,
	).Scan(&id)
	if err != nil {
		return err
	}

	for i, destination := range realEstate.Destinations {
		const statement = `
INSERT INTO real_estate_destinations
  (real_estate_id, num, destination)
VALUES ($1, $2, $3)
ON CONFLICT (real_estate_id, num)
  DO UPDATE SET destination=EXCLUDED.destination
`
		_, err := r.db.Exec(statement, id, i, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Database) ListRealEstates(
	longitude float32,
	latitude float32,
	distance float32,
) ([]RealEstate, error) {
	const query = `
SELECT
   real_estate_id,
   address       ,
   occurance_type,
   area          ,
   initial_price ,
   district      ,
   longitude     ,
   latitude      
FROM real_estates
WHERE acos(
    sin(radians(latitude))*sin(radians($2))
    +cos(radians(latitude))*cos(radians($2))*cos(radians($1-longitude))
  )*6371 < $3
`
	rows, err := r.db.Query(query, longitude, latitude, distance)
	if err != nil {
		return nil, fmt.Errorf("running query: %w", err)
	}
	defer rows.Close()

	var result []RealEstate
	for rows.Next() {
		var realEstate RealEstate

		err = rows.Scan(
			&realEstate.Id,
			&realEstate.Address,
			&realEstate.OccuanceType,
			&realEstate.Area,
			&realEstate.InitialPrice,
			&realEstate.District,
			&realEstate.Location.Longitude,
			&realEstate.Location.Latitude,
		)
		if err != nil {
			slog.Error("scanning row", err)
		}

		err = binding.Validator.ValidateStruct(realEstate)
		if err != nil {
			slog.Error("validating", "error", err)
			continue
		}

		result = append(result, realEstate)
	}

	return result, nil
}

func (r *Database) ListPolygons() ([]Polygon, error) {

	const query = `
WITH warsaw_boundary AS (
    SELECT
        way AS geom
    FROM
        planet_osm_polygon
    WHERE
        name = 'Warszawa'
        AND boundary = 'administrative'
        AND admin_level = '6'
),
vertices AS (
    SELECT
        osm_id,
        name,
        ST_Transform((ST_DumpPoints(way)).geom, 4326) AS geom_4326  -- Convert to EPSG:4326 geometry
    FROM
        planet_osm_polygon p, warsaw_boundary w
    WHERE
        p.boundary = 'administrative'
        AND p.admin_level = '8'
        AND ST_Intersects(p.way, w.geom)
)
SELECT
    osm_id,
    name,
    ST_X(geom_4326) AS longitude,
    ST_Y(geom_4326) AS latitude
FROM
    vertices;
`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("running query: %w", err)
	}
	defer rows.Close()

	result := make(map[string]Polygon)
	for rows.Next() {
		var id int
		var name string
		var longitude float32
		var latitude float32

		err = rows.Scan(&id, &name, &longitude, &latitude)
		if err != nil {
			slog.Error("scanning row", err)
		}

		location := Location{Longitude: longitude, Latitude: latitude}

		old, exists := result[name]
		if exists {
			old.Locations = append(old.Locations, location)
			result[name] = old
		} else {
			polygon := Polygon{
				Id:        id,
				Name:      name,
				Locations: []Location{location},
			}
			result[name] = polygon
		}
	}

	return maps.Values(result), nil
}
