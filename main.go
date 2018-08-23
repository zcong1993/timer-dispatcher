package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/zcong1993/timer-dispatcher/core/tasks"
	"github.com/zcong1993/timer-dispatcher/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// PdService is our grpc service
type PdService struct {
	scheduler *tasks.Scheduler
}

// Add implement PdService
func (pd *PdService) Add(ctx context.Context, in *pb.Task) (*pb.Resp, error) {
	pd.scheduler.Push(in.Value, in.Timestamp)
	return &pb.Resp{Message: "ok", Ok: true}, nil
}

func runRpcServer(port string) {
	s := tasks.NewScheduler(time.Second*5, 10)
	ss := grpc.NewServer()
	pb.RegisterTdServiceServer(ss, &PdService{scheduler: s})

	ch := s.MsgCh()

	go func() {
		for msg := range ch {
			println(msg.Timestamp, msg.Value)
		}
	}()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	if err = ss.Serve(listener); err != nil {
		log.Fatal("ListenTCP error:", err)
	}
}

func runGatewayServer(rpcPort, port string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterTdServiceHandlerFromEndpoint(ctx, mux, rpcPort, opts)

	if err != nil {
		log.Fatal("Serve http error:", err)
	}

	http.ListenAndServe(port, mux)
}

func main() {
	if os.Getenv("GATEWAY") == "true" {
		println("run gateway server on :8080")
		go runGatewayServer(":1234", ":8080")
		println(time.Now().UnixNano())
	}
	runRpcServer(":1234")
}
