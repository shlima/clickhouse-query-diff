FROM node:24-alpine3.21 AS frontend
ARG GIT_COMMIT
ARG VERSION
ENV NUXT_PUBLIC_API_PATH=""
ENV NUXT_PUBLIC_VERSION=${VERSION}
WORKDIR /app
COPY frontend/ /app
RUN yarn
RUN yarn generate

FROM golang:1.24 AS builder
ARG LDFLAGS
ARG GIT_COMMIT
ARG VERSION
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY main.go .
RUN go mod download
RUN CGO_ENABLED=0 go build -o build/clickhouse-query-diff -ldflags="${LDFLAGS}" main.go

FROM debian:bookworm-slim
ARG GIT_COMMIT
ARG VERSION
WORKDIR /app
COPY --from=frontend /app/.output /app/frontend/.output
COPY --from=builder /app/build/clickhouse-query-diff /usr/local/bin/clickhouse-query-diff
EXPOSE 33333
ENTRYPOINT ["/usr/local/bin/clickhouse-query-diff"]
