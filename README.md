<h1 align="center">
  <img src="https://i.ibb.co/LhGDRPq/Solution-5-File-Metadata.png" alt="Logo" width="600">
</h1>

## Features

- Register the user
- Add exercise to the user
- List of exercises from a user with or without filter (from (date), to (date), and limit (number))

## Tech Stack

**Server:** Golang (Gin) and Mongo

## Run Locally

**Your server is accessible in port 3000**
Here is the steps to run it with `golang`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-exercise-tracker-ms.git

# Move to project
$ cd fcc-exercise-tracker-ms

# Set gin to release mode
$ export GIN_MODE=release

# Run the application
$ go run main.go
```

Here is the steps to run it with `docker-compose`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-exercise-tracker-ms.git

# Move to project
$ cd fcc-exercise-tracker-ms

# Set gin to release mode
$ export GIN_MODE=release

# Download, setup, and run the image
$ docker-compose up -d

# Stops containers and removes containers, networks, volumes, and images created by up
$ docker-compose down
```

## Running Tests

To run tests and get the percentage of code coverage, run the following command

```bash
  go test ./... -cover
```

## Lessons Learned

- How to create API with Golang (Gin)
- How to create middleware
- How to create a gin handler
- How to create a unit test for the handler layer
- How to create a service layer
- How to create a unit test for the service layer
- How to create a repository layer
- How to connect to the mongo database
- How to create an image from a service
- How to setup docker container to wrap all needed dependencies for a service

## API Reference

[Test API with REST Client Extension](https://github.com/madjiebimaa/fcc-exercise-tracker-ms/tree/main/docs/apis)
