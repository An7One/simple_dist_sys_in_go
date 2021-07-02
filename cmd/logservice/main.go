package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/an7one/tutorial/simple_dist_sys_in_go/log"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/registry"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/service"
)

func main() {
	log.Run("./distributed.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stdlog.Fatalln(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down the service")
}
