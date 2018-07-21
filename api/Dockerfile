FROM golang:1.10-alpine AS build-env

RUN apk --update add make git

WORKDIR /go/src/github.com/fujimisakari/grpc-study/api
ADD . /go/src/github.com/fujimisakari/grpc-study/api
RUN go get github.com/golang/dep/cmd/dep && \
    dep ensure -v && \
    go build -o bin/api

FROM alpine:latest

COPY --from=build-env /go/src/github.com/fujimisakari/grpc-study/api/bin/api /bin/api

RUN apk add --no-cache tzdata ca-certificates

CMD ["api"]
