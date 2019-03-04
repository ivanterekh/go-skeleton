FROM golang:alpine AS builder

WORKDIR $GOPATH/src/github.com/ivanterekh/go-skeleton/
COPY . .
RUN apk update && apk add --no-cache ca-certificates git make tzdata \
    && adduser -D -g '' appuser \
    && update-ca-certificates \
    && APP=/go-skeleton/app make build

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /go-skeleton/
COPY --from=builder /go-skeleton/app /go-skeleton/app
COPY config.json .
EXPOSE 8080
USER appuser

ENTRYPOINT ["./app"]

