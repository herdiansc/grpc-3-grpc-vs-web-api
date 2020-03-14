package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-3-grpc-vs-web-api/proto"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err.Error())
	}
	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	n, _ := strconv.Atoi(os.Args[1])
	start := time.Now()
	for i := 0; i < n; i++ {
		r, err := c.SayHello(ctx, &pb.HelloWorldRequest{Name:"world"})
		if err != nil {
			log.Fatalf("failed get merchant: %v", err.Error())
		}
		log.Printf("get merchant: %v", r)
	}
	ms := time.Since(start)
	log.Printf("mil: %v\tmic: %v\tmin:%v\n", ms.Milliseconds(), ms.Microseconds(), ms.Minutes())
	log.Printf("total time: %v\ttotal req: %v\ttime per req:%v\n", ms, n, ms.Microseconds()/int64(n))
}
