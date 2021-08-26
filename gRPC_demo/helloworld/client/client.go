package main

import (
	"context"
	"fmt"
	pb "gPRC_demo/helloworld/pd"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8972",grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect:%v",err)
	}
	defer  conn.Close()

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(),&pb.HelloRequest{Name:"Gentle"})
	if err != nil {
		fmt.Printf("could not greet:%v",err)
	}
	fmt.Printf("Greeting:%s!\n",r.Message)
}
