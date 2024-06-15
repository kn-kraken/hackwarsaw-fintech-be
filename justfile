api:
  go run ./cmd/api

watch:
  watchexec -re go just api

scrap:
  go run ./cmd/scrapper
