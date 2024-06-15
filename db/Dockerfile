# Use the official PostGIS image from the Docker Hub
FROM postgis/postgis:latest

# Install nano text editor
RUN apt-get update && \
    apt-get install -y nano && \
    rm -rf /var/lib/apt/lists/*

# Install necessary packages for downloading OSM data
RUN apt-get update && \
    apt-get install -y wget osm2pgsql && \
    rm -rf /var/lib/apt/lists/*

# Set environment variables
ENV POSTGRES_DB=gis
ENV POSTGRES_USER=gisuser
ENV POSTGRES_PASSWORD=gispassword

# Create a directory for OSM data
RUN mkdir -p /osmdata

# Copy the OSM data file and the script to load OSM data
COPY mazowieckie-latest.osm.pbf /osmdata/
# COPY load_osm.sh /docker-entrypoint-initdb.d/

# Make the script executable
# RUN chmod +x /docker-entrypoint-initdb.d/load_osm.sh

# Expose the default PostgreSQL port
EXPOSE 5432
