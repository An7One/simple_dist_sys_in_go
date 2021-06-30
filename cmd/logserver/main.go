package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/an7one/tutorial/simple_dist_sys_in_go/log"
	"github.com/an7one/tutorial/simple_dist_sys_in_go/service"
)

func main() {
	log.Run("./distributed.log")

	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)

	if err != nil {
		stdlog.Fatalln(err)
	}

	<-ctx.Done()

	fmt.Println("Shutting down the service")
}
