package main

import (
	"github.com/WiFeng/go-sky"
	"github.com/WiFeng/go-sky-helloworld/pkg/endpoint"
	"github.com/WiFeng/go-sky-helloworld/pkg/service"
	"github.com/WiFeng/go-sky-helloworld/pkg/transport/http"
)

func main() {

	var (
		service     = service.New()
		endpoints   = endpoint.New(service)
		httpHandler = http.NewHandler(endpoints)
	)

	sky.Run(httpHandler)
}
