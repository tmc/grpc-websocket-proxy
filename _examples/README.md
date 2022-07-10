## wsechoserver
This example implements a websocket server using the proxy. The proto definition contains three RPCs:
- `Echo` is a bidirectional stream that echoes back the client request.
- `Stream` is a server-side stream that sends a fixed number of messages to the client and closes the connection
- `Heartbeat` is a bidirectional stream that sends back messages to the client at intervals.

## Usage
Ideally, you should be able to run `go run _examples/cmd/wsechoserver` and start the example server.
To build it locally, you'll need to install:
- [buf](https://buf.build/)
- [protoc-gen-go](google.golang.org/protobuf/cmd/protoc-gen-go)
- [protoc-gen-go-grpc](google.golang.org/grpc/cmd/protoc-gen-go-grpc)
- [protoc-gen-go-grpc-gateway](github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway)

Generate the protobuf files by running:
`buf generate -v .`

You should now be able to start the server with the same run command above.