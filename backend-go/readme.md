# backend-go

A Go backend. Technologies used:

- Chi (routing)
- Postgres (persistent data)
- Redis (caching)
- NewRelic (logging)

# Development

Developed using Go version 1.22. Older versions may or may not work, depending on the dependencies used in this code.

To download dependencies, run `go mod tidy`.

# Running on local environment

It is preferrable to use Docker, because the Dockerfile includes hot reload, and environment variables. Before starting the containers, edit the `NR_LICENSE` environment variable in `docker-compose.yml` to be your own NewRelic license key.