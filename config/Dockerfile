FROM golang:1.24.3 AS builder

WORKDIR /usr/src/app

COPY ../go.mod ../go.sum ./

COPY ../cmd ./cmd
COPY ./pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/API ./cmd

FROM scratch

COPY --from=builder /app/API /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /

CMD ["/API"]
