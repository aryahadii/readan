FROM golang:1.11.0 as builder
WORKDIR /go/src/github.com/aryahadii/readan
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make readan

FROM ubuntu:18.04
WORKDIR /root/
COPY --from=builder /go/src/github.com/aryahadii/readan/readan .
ENTRYPOINT ["./readan"]
