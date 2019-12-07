# build stage
FROM golang:buster as builder
RUN apt-get install ca-certificates
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=${GOOS} go build -a -installsuffix cgo -o main .


# final stage
FROM scratch as production
COPY --from=builder /app/main /fritzbox-graphite
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
EXPOSE 8000
ENTRYPOINT ["/fritzbox-graphite"]