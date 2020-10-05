package main

import (
	"context"
	"go-learn-project/01_middleware_sdk/15_rpc_framework/grpc_etcd_gin/proto"
	"go-learn-project/01_middleware_sdk/15_rpc_framework/grpc_etcd_gin/registercenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Greeter  struct{}

func (greeter *Greeter) Helloservice(ctx context.Context, req *protocol.RequestMessage) (*protocol.ResponseMessage, error) {
	log.Print("received ")
	responseMessage:= protocol.ResponseMessage{Msg: "hello "+req.Name}
	return &responseMessage,nil
}
func main() {
	lis,err:=net.Listen("tcp","10.3.20.223:8084")
	////lis,err:=net.Listen("tcp","127.0.0.1:8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go registercenter.Register("192.168.172.5:2379;192.168.172.6:2379;192.168.172.7:2379","GreeterService","10.3.20.223:8084",10)
	//go registercenter.Register("127.0.0.1:2379","GreeterService","127.0.0.1:8084",10)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		registercenter.UnRegister("GreeterService", "10.3.20.223:8084")
		//registercenter.UnRegister("GreeterService", "127.0.0.1:8084")

		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()
	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}
	s:=grpc.NewServer(grpc.MaxSendMsgSize(10*1024*1024),
		grpc.MaxRecvMsgSize(10*1024*1024),
		grpc.KeepaliveParams(kasp));
	greeter:=Greeter{}
	protocol.RegisterGreeterServer(s,&greeter)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

