FROM golang:1.19-alpine AS builder

WORKDIR /go/src/dev

RUN apk add tzdata

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -o ../app/app main.go

RUN rm -rf /go/src/dev


FROM alpine:latest AS runtime

WORKDIR /go/src/app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/src/app/app ./app

COPY config.yaml .

EXPOSE 8080

CMD ["./app", "http"]
