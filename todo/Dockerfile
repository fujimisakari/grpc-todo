FROM golang:1.10-stretch

# install protoc
# RUN apt-get update && apt-get -y install unzip && apt-get clean
# ENV PB_VER 3.5.1
# ENV PB_URL https://github.com/google/protobuf/releases/download/v${PB_VER}/protoc-${PB_VER}-linux-x86_64.zip
# RUN mkdir /protoc && \
#     curl -L ${PB_URL} > /protoc/protoc.zip && \
#     cd /protoc && \
#     unzip protoc.zip
# RUN go get -u github.com/golang/protobuf/protoc-gen-go

Run go get github.com/fujimisakari/grpc-study/todo
WORKDIR /go/src/github.com/fujimisakari/grpc-study/todo
RUN go build

CMD ["todo"]
