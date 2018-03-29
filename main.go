package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"db-srv/handler"
	"db-srv/subscriber"

	example "db-srv/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.weipingyun.db.srv.db-srv"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.com.weipingyun.db.srv.db-srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.com.weipingyun.db.srv.db-srv", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
