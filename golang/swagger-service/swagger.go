package main

import (
	"github.com/ic2hrmk/snips/golang/swagger-service/resource"
	"github.com/ic2hrmk/snips/golang/swagger-service/service"
	"log"
)

const (
	defaultServicePort = ":8080"
)

func main() {
	baseService := service.NewBaseService()
	baseResource := resource.NewBaseResourceWithSwagger(baseService)

	log.Printf("swagger documentation is available at [%s]\n", defaultServicePort)

	if err := baseResource.RunStandalone(defaultServicePort); err != nil {
		log.Fatalf("failed to [serve] Swagger service, %s", err.Error())
	}
}
