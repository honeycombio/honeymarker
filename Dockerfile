FROM golang:1.18.1-alpine3.15 as builder

RUN apk update

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

FROM scratch

COPY --from=builder /app/honeymarker /usr/bin/honeymarker
