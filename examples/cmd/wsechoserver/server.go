package main

import (
	"fmt"
	"time"

	"github.com/tmc/grpc-websocket-proxy/examples/cmd/wsechoserver/echoserver"
)

type Server struct{}

func (s *Server) Stream(_ *echoserver.Void, stream echoserver.EchoService_StreamServer) error {
	start := time.Now()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		if err := stream.Send(&echoserver.EchoResponse{
			Message: "hello there!" + fmt.Sprint(time.Now().Sub(start)),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) Echo(srv echoserver.EchoService_EchoServer) error {
	for {
		req, err := srv.Recv()
		if err != nil {
			return err
		}
		if err := srv.Send(&echoserver.EchoResponse{
			Message: req.Message + "!",
		}); err != nil {
			return err
		}
	}
}
