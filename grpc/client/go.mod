module github.com/renuka-fernando/examples/grpc/client

go 1.22.2

require (
	github.com/renuka-fernando/examples/grpc/server v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.64.0
)

require (
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/renuka-fernando/examples/grpc/server => ../server
