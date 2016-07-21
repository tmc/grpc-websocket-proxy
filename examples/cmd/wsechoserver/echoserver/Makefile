all:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		echoserver.proto
