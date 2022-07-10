#!/bin/bash
# Start test db container
docker-compose up -d
docker-compose port db 5432