FROM golang:latest AS builder

WORKDIR /

COPY ./ /

RUN go mod download
RUN go build -mod=readonly -v -o /tmp/application

FROM ubuntu:latest

COPY --from=builder /tmp/application /


ENTRYPOINT ["/application"]