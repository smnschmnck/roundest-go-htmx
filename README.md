# 🐉 Roundest Mon

Let users vote on which Pokémon they think is the roundest to find out what the roundest Pokémon is.

This project is inspired by the [1 App, 5 Stacks](https://github.com/t3dotgg/1app5stacks) project by @t3dotgg which implements [Roundest Mon](https://github.com/t3dotgg/roundest-mon) in 5 different tech stacks. I felt like a Go + HTMX version was missing so I built it.

## Tech Stack

- [Go](https://github.com/golang/go)
- [gomponents](https://github.com/maragudk/gomponents)
- [HTMX](https://github.com/bigskysoftware/htmx)
- [Tailwind CSS](https://github.com/tailwindlabs/tailwindcss)
- [Labstack Echo](https://github.com/labstack/echo)
- [sqlc](https://github.com/sqlc-dev/sqlc)
- [PostgreSQL](https://www.postgresql.org/)

## Running Locally

### Prerequisites

- Go installed
- Node.js installed
- pnpm installed
- PostgreSQL running somewhere

### Steps

1. Migrate database using schema in `db/sql/schema.sql`
1. Add `.env` file with variables `DATABASE_URL` and `PORT`
1. Install dependencies with `pnpm install`
1. Generate CSS with `npm run tailwind:watch`
1. Run app with `go run main.go`

## Building the App

### Steps

1. Migrate database using schema in `db/sql/schema.sql`
1. Add `.env` file with variables `DATABASE_URL` and `PORT`
1. Install dependencies with `pnpm install`
1. Build app with command `npm run build`
