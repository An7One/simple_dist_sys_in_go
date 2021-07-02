package main

import (
	"context"
	"fmt"
	stdLog "log"

	"github.com/an7one/tutorial/simple_dist_sys_in_go/grade"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/log"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/registry"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddress + "/serivces",
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grade.RegisterHandlers)
	if err != nil {
		stdLog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
