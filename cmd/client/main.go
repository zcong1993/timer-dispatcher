package main

import (
	"context"
	"fmt"
	"github.com/zcong1993/timer-dispatcher/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func makeAdd(c pb.TdServiceClient, value string, timestamp int64) {
	resp, err := c.Add(context.Background(), &pb.Task{Value: value, Timestamp: timestamp})
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}

	fmt.Printf("%t - %s\n", resp.Ok, resp.Message)
}

func main() {
	c, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	client := pb.NewTdServiceClient(c)

	now := time.Now()

	for i := 0; i < 100; i++ {
		makeAdd(client, fmt.Sprintf("test - %d", i), now.Add(time.Duration(int32(i))*time.Second).UnixNano())
	}
}
