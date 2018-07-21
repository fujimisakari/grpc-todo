FROM golang:1.10-alpine AS build-env

RUN apk --update add make git

WORKDIR /go/src/github.com/fujimisakari/grpc-study/todo
ADD . /go/src/github.com/fujimisakari/grpc-study/todo
RUN go get github.com/golang/dep/cmd/dep && \
    dep ensure -v && \
    go build -o bin/todo

FROM alpine:latest

COPY --from=build-env /go/src/github.com/fujimisakari/grpc-study/todo/bin/todo /bin/todo

RUN apk add --no-cache tzdata ca-certificates

CMD ["todo"]
