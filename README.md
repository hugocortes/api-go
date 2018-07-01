# paprika-api
simple golang api

# Running
1. Create `.env` using `.env.example`
2. `go get ./`
3. `go run main.go`

# Building
1. Create `.env` using `.env.example`
2. `go get ./`
3. `go build -o paprika-api ./`
4. `./paprika-api`

# Testing
1. Create `.env` using `.env.example`
2. `go get ./`
3. `go test ./test/*`
  * Tests are modularized to test individual components of the api, i.e. server, controllers, db, etc.
