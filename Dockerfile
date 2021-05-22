
FROM alpine:3.13.1 AS base
EXPOSE 4000

FROM golang:1.16.3-alpine AS builder
RUN apk update
RUN apk add build-base
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o wallserver -ldflags "-s" main.go

FROM base as FINAL
WORKDIR /app
COPY --from=builder /build/wallserver .
CMD [ "/app/wallserver" ]