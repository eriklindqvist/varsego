#!/bin/sh

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
docker build -t varsego .
docker tag varsego:latest proto:5000/varsego
