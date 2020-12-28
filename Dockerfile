FROM golang:1.15.6-buster as builder
WORKDIR /go/src/gitlab.jave.fr/tf03kmon
COPY *.go *.mod /go/src/gitlab.jave.fr/tf03kmon/
RUN go build

FROM debian:buster
WORKDIR /root/
COPY --from=builder /go/src/gitlab.jave.fr/tf03kmon /root/
COPY config.yaml .
CMD ["/root/tf03kmon", "-config=/root/config.yaml" , "-loglevel=debug" ]
