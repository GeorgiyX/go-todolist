# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go get -u github.com/mailru/easyjson/...
ENV PATH=$GOPATH/bin:$PATH

COPY . .
RUN easyjson -all -pkg app/models
RUN go build -o api.run ./cmd/.

# Deploy stage
FROM ubuntu

ENV ARTIFACT api.run
ENV API_PORT 5000

WORKDIR /app
COPY --from=build /app/$ARTIFACT $ARTIFACT

ENV MODE release

CMD ./$ARTIFACT

EXPOSE $API_PORT
