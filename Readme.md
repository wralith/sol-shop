# Sol Shop

This is an example Monolith application with few basic concepts of an e-commerce system.

### Dev Setup

```sh
cd server
make gen-ent
go run cmd/main.go
```

REST docs will start with app at docs endpoint: `http://localhost:8080/docs` thanks to [swaggo/swag](https://github.com/swaggo/swag) and [rapidoc](https://github.com/rapi-doc/RapiDoc)

# Stack

Backend written in `go` and mainly shaped around `echo` and `ent` frameworks.

# Todos

- Request and Response bodies? Separate or same, custom or as ent generated? Validations?
- Docs
- Error Handling
- Tests!
- Logging
