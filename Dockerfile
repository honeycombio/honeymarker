FROM golang:1.16.5-alpine3.13

RUN apk update && rm -rf /var/cache/apk/*

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download
RUN go mod verify

ADD . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
    -o honeymarker

RUN mv /app/honeymarker /usr/bin/honeymarker
