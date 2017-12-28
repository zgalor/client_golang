# This Dockerfile builds an image for a client_golang example.
#
# Use as (from the root for the client_golang repository):
#    docker build -f examples/$name/Dockerfile -t prometheus/golang-example-$name .

# Builder image, where we build the example.
FROM golang:1.9.0
WORKDIR /$GOPATH/src/github.com/prometheus/client_golang
COPY . .
WORKDIR /$GOPATH/src/github.com/prometheus/client_golang/prometheus
RUN go get -d
WORKDIR /$GOPATH/src/github.com/prometheus/client_golang/examples/reload
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'

# Final image.
FROM scratch
LABEL maintainer "The Prometheus Authors <prometheus-developers@googlegroups.com>"
COPY --from=builder /$GOPATH/src/github.com/prometheus/client_golang/examples/reload .
EXPOSE 8080
ENTRYPOINT ["/reload"]
