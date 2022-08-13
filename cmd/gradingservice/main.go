package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/mosesbenjamin/distributed_system/grades"
	"github.com/mosesbenjamin/distributed_system/registry"
	"github.com/mosesbenjamin/distributed_system/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
