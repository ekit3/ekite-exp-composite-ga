
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod ./
COPY entrypoint.sh /cs-entrypoint.sh

RUN chmod +x /cs-entrypoint.sh
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go

FROM alpine:3.10 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-go /docker-go
COPY --from=build-stage --chmod=755 /cs-entrypoint.sh /cs-entrypoint.sh

ENTRYPOINT ["/cs-entrypoint.sh"]
