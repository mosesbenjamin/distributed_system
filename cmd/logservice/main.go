package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/mosesbenjamin/distributed_system/log"
	"github.com/mosesbenjamin/distributed_system/registry"
	"github.com/mosesbenjamin/distributed_system/service"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")

}
