FROM golang:1.16.6-alpine as builder

WORKDIR /code/
COPY . ./

RUN apk add git build-base
RUN go build -trimpath -ldflags="-s -w" -o /app cmd/bot/main.go

FROM alpine:3.15 as alpine
RUN apk add --no-cache git ca-certificates tzdata
COPY --from=builder app .

CMD ["./app"]
