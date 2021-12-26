# pull official base image
FROM postgres:12.2-alpine

# run init.sql
ADD init.sql /docker-entrypoint-initdb.d


# FROM golang:latest

# RUN mkdir /app
# WORKDIR /app/server

# COPY go.mod ./src/main
# COPY go.sum ./src/main
# RUN go mod download
# COPY . .