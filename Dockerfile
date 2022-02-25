FROM golang:1.17-alpine as builder
RUN apk --update add ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /polygon-otel-collector

FROM scratch

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /polygon-otel-collector /
COPY polygon-config.yaml /etc/otel/config.yaml

EXPOSE 4317 55680 55679 8086

ENTRYPOINT ["/polygon-otel-collector"]
CMD ["--config", "/etc/otel/config.yaml"]
