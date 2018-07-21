FROM golang:1.10-alpine AS build-env

RUN apk --update add make git

WORKDIR /go/src/github.com/fujimisakari/grpc-study/dashboard
ADD . /go/src/github.com/fujimisakari/grpc-study/dashboard
RUN go get github.com/golang/dep/cmd/dep && \
    dep ensure -v && \
    go build -o bin/dashboard

FROM alpine:latest

COPY --from=build-env /go/src/github.com/fujimisakari/grpc-study/dashboard/bin/dashboard /bin/dashboard

RUN apk add --no-cache tzdata ca-certificates

CMD ["dashboard"]
