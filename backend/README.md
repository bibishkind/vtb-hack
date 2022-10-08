# Backend app for VTB Hack

## Build & Run (locally)
### Prerequisites
- go 1.16
- docker & docker compose

Create .env in root directory and add following values:
```dotenv
DB_URL=postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
SIGNING_KEY=secret
```

Use `docker-compose up` to build and run project.