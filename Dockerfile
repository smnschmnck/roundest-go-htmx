FROM node:22-bookworm AS node

FROM golang:1.24.3-bookworm AS builder

WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=node /usr/local/ /usr/local/

RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile --config.strict-dep-builds=false

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN pnpm run build

FROM debian:bookworm-slim AS runtime

WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && useradd --system --create-home --uid 1001 appuser \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main ./main
COPY --from=builder /app/.build ./.build

ENV PORT=3000

USER appuser

EXPOSE 3000

CMD ["./main"]
