# API Development Notes

## Endpoints

API Endpoints are currently each stored in a unique file prefixed with `ep_`.

Each Endpoint is assigned a function and the function is assigned to a route.

Routes are configured within root.go as is the gin-gonic http server.

## HTTP Server

Is currently using [gin-gonic](https://pkg.go.dev/github.com/gin-gonic/gin).

The HTTP Server is being configured within the `shifter/api/root.go` folder.


## API Struct Objects
All Strcuts relating to API Endpoints and responses are stored within the `structs.go` file located at `shifter/api/structs.go`
