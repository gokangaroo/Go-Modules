package main

import (
	"context"
	"google.golang.org/grpc"
	"hello/go_client/proto/hello"
	"io"
	"log"
)

const (
	// grpc 服务地址
	Address = "0.0.0.0:9090"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//初始化客户端
	c := hello.NewHelloClient(conn)

	//调用SayHello 方法
	res, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)

	//调用 LotsOfReplies 方法
	stream, err := c.LotsOfReplies(context.Background(), &hello.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv: %v", err)
		}
		log.Printf("%s", res.Message)
	}
}
