package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
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

func (r *Database) ListBusinessesInArea() ([]Business, error) {
	const query = `
SELECT
    name,
    ST_X(ST_Transform(way, 4326)) AS lon,
    ST_Y(ST_Transform(way, 4326)) AS lat,
    ST_Distance(
        ST_SetSRID(ST_MakePoint(21.045535, 52.256403), 4326),
        ST_Transform(way, 4326)
    ) AS distance
FROM planet_osm_point
WHERE
    ST_DWithin(
        ST_SetSRID(ST_MakePoint(21.045535, 52.256403), 4326),
        ST_Transform(way, 4326),
        1000
    )
    AND name != ''
    AND shop = 'greengrocer'
ORDER BY distance;
`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("running query: %w", err)
	}
	defer rows.Close()

	var result []Business
	for rows.Next() {
		var business Business

		err = rows.Scan(&business.Name,
			&business.Location.Longitude,
			&business.Location.Latitude,
			&business.Distance,
		)
		if err != nil {
			slog.Error("scanning row", err)
		}

		result = append(result, business)
	}
	return result, nil
}
