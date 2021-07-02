package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/an7one/tutorial/simple_dist_sys_in_go/log"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/portal"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/registry"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/service"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stdlog.Fatal(err)
	}
	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.PortalService,
		ServiceURL:  serviceAddress,
		RequiredServices: []registry.ServiceName{
			registry.LogService,
			registry.GradingService,
		},
		ServiceUpdateURL: serviceAddress + "/services",
	}

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		portal.RegisterHandlers)
	if err != nil {
		stdlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err != nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down portal.")
}
