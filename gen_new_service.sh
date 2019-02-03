export PATH=$PATH:$PWD/bin
cd $PWD/src/events/
protoc -I events/ events/events.proto --go_out=plugins=grpc:events