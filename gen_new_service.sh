export PATH=$PATH:$PWD/bin
cd $PWD/src/vaults/
protoc -I vaults/ vaults/vaults.proto --go_out=plugins=grpc:vaults