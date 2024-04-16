FROM golang:1.22.0-alpine3.18 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.22.0-alpine3.18 as builder
RUN apk add --no-cache git make gcc musl-dev libc6-compat
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 \
    go build -buildvcs=false -o /bin/app ./cmd/app

# Step 3: Final
FROM alpine:3.18
COPY --from=builder /app/config /config
# COPY --from=builder /app/migrations /migrations
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app"]