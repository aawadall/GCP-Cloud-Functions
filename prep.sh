#if [ $GOPATH -eq "" ]
#then
    echo "Fixing GOPATH"
    export GOPATH=$PWD/gopath
#fi
go get -u google.golang.org/grpc
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0rc1/protoc-3.7.0-rc1-linux-x86_64.zip
unzip ./protoc-3.7.0-rc1-linux-x86_64.zip
export PATH=$PATH:$PWD/bin
go get -u github.com/golang/protobuf/protoc-gen-go