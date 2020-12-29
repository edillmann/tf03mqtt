FROM golang:1.15.6-buster as builder
WORKDIR /go/src/github.com/edillmann/tf03mqtt
COPY *.go *.mod /go/src/github.com/edillmann/tf03mqtt/
RUN go build

FROM debian:buster
WORKDIR /root/
COPY --from=builder /go/src/github.com/edillmann/tf03mqtt /root/
COPY config.yaml .
CMD ["/root/tf03mqtt", "-config=/root/config.yaml" , "-loglevel=debug" ]
