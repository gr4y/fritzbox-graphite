# Build Stage
FROM golang:buster as build-env
RUN apt-get install ca-certificates
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /app

# Final Stage
FROM scratch
COPY docker /
COPY --from=build-env /etc/ssl/certs /etc/ssl/certs
COPY --from=build-env /app/fritzbox-graphite /app/fritzbox-graphite
ENTRYPOINT ["/app/fritzbox-graphite"]