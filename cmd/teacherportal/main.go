package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/mosesbenjamin/distributed_system/log"
	"github.com/mosesbenjamin/distributed_system/portal"
	"github.com/mosesbenjamin/distributed_system/registry"
	"github.com/mosesbenjamin/distributed_system/service"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.TeacherPortal
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		portal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down teacher portal")

}
