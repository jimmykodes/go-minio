# BUILD
FROM golang:1.14-buster AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -o /app ./cmd/main

# RUN
FROM debian:buster
RUN apt-get update --fix-missing && \
    apt-get install -yqq ca-certificates

COPY --from=builder /app .
CMD ["./app"]
