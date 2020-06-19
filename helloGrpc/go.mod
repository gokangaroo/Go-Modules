module hello

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v0.0.0-00010101000000-000000000000
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.22.1
