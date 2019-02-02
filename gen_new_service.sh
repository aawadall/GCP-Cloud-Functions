export PATH=$PATH:$PWD/bin
cd $PWD/src/google.golang.org/grpc/examples/helloworld/
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld