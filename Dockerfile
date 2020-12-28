FROM golang:1.15.6-buster as builder
WORKDIR /go/src/github.com/edillmann/tf03kmon
COPY *.go *.mod /go/src/github.com/edillmann/tf03kmon/
RUN go build

FROM debian:buster
WORKDIR /root/
COPY --from=builder /go/src/github.com/edillmann/tf03kmon /root/
COPY config.yaml .
CMD ["/root/tf03kmon", "-config=/root/config.yaml" , "-loglevel=debug" ]