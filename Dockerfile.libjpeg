FROM golang:1.18 as builder

RUN apt update && \
    apt install -y libjpeg-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app libjpeg.go

ENTRYPOINT ["/app/app"]