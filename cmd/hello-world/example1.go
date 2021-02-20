package main

import (
	"context"
	greeter "github.com/WalkerLiuFei/micro-examples/proto/greeter"
	"github.com/micro/go-micro"
	"log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	rsp.Greeting = "Hello" + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
	)

	service.Init()

	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
