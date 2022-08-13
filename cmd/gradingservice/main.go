package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/mosesbenjamin/distributed_system/grades"
	"github.com/mosesbenjamin/distributed_system/log"
	"github.com/mosesbenjamin/distributed_system/registry"
	"github.com/mosesbenjamin/distributed_system/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	} else {
		stlog.Println(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
