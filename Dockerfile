FROM golang

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobug/protoc-gen-go
RUN export PATH=$PATH:$GOPATH/bin
WORKDIR $GOPATH/src/google.golang.org/grpc/examples/helloworld
RUN go run greeter_server/main.go
CMD go run greeter_client/main.go