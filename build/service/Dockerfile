FROM golang:1.20-alpine3.17 AS builder

# RUN apk add --no-cache tesseract-ocr-dev tesseract-ocr-data-rus g++

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./.bin/app ./cmd/subscription-service/main.go

FROM alpine:3.17

# RUN apk add --no-cache tesseract-ocr-dev tesseract-ocr-data-rus 

WORKDIR /app

COPY --from=builder /app/.bin/app .

COPY ./config ./config

CMD ./app
