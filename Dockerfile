# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /opt
COPY ampl.tgz ./
RUN tar -xzf ampl.tgz
ENV PATH="${PATH}:/opt/ampl_linux-intel64"

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/hackwarsaw-fintech-be ./cmd/api 

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/hackwarsaw-fintech-scrapper ./cmd/scrapper 

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
ENTRYPOINT sleep 10 && \
  /app/hackwarsaw-fintech-be \
  -dbhost "$DB_HOST" \
  -dbport "$DB_PORT" \
  -dbname "$DB_NAME" \
  -dbuser "$DB_USER" \
  -dbpassword "$DB_PASSWORD"
  #  & \
  # /app/hackwarsaw-fintech-scrapper \
  #   -dbhost "$DB_HOST" \
  #   -dbport "$DB_PORT" \
  #   -dbname "$DB_NAME" \
  #   -dbuser "$DB_USER" \
  #   -dbpassword "$DB_PASSWORD" \
  #   -geocoding-apikey "$GEOCODING_APIKEY"
