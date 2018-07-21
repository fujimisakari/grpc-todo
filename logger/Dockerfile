FROM golang:1.10-alpine AS build-env

RUN apk --update add make git

WORKDIR /go/src/github.com/fujimisakari/grpc-study/logger
ADD . /go/src/github.com/fujimisakari/grpc-study/logger
RUN go get github.com/golang/dep/cmd/dep && \
    dep ensure -v && \
    go build -o bin/logger

FROM alpine:latest

COPY --from=build-env /go/src/github.com/fujimisakari/grpc-study/logger/bin/logger /bin/logger

RUN apk add --no-cache tzdata ca-certificates

CMD ["logger"]
