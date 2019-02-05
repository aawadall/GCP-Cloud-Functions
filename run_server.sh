export PATH=$PATH:$PWD/bin
echo "Running server"

go run src/google.golang.org/grpc/examples/helloworld/greeter_server/main.go