# Fintech BE

## Depencencies

- `go 1.22.4`
- `ampl`
- Google's geocoding API key
- `docker`
- `docker-compose`
- `just`

## Run with docker

```sh
docker-compose up
```

API will be available at `http://localhost:8080`.

## Scrapping data

Initially the database will be empty. Scrapper needs to be run in order to
populate it with scrapped data.

```sh
GEOCODING_APIKEY="..." just scrap
```

## Deploy link

`https://kraken.anczykowski.com/fintech/`

## Example requests

- `GET /real-estates?business=ALCOHOL&latitude=52.256403&longitude=21.0425535&range=1`
  -- get real estates and businesses for category alcohol within the given vicinity
- `GET /polygons` -- get areas of interest
