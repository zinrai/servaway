FROM golang:1.24-alpine AS build

WORKDIR /src
COPY go.mod main.go ./
RUN go build -o /servaway

FROM alpine:3.21

COPY --from=build /servaway /usr/local/bin/servaway
CMD ["servaway"]
