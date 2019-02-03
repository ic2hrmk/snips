package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"log"
	"net/http"
)

const (
	defaultServicePort = ":8080"
)

func Factory() (http.Handler, error) {
	var webServices []*restful.WebService
	container := restful.NewContainer()

	c := restfulspec.Config{
		WebServices:    webServices,
		WebServicesURL: "https://my-project.com/",
		APIPath:        "/swagger/apidocs.json",
		PostBuildSwaggerObjectHandler: func(swo *spec.Swagger) {
			swo.Info = &spec.Info{
				InfoProps: spec.InfoProps{
					Title: "My project API docs",
					Contact: &spec.ContactInfo{
						Name: "my-project Team",
					},
					Version: "v0.1.0",
				},
			}
		},
	}

	container.Add(restfulspec.NewOpenAPIService(c))

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		CookiesAllowed: false,
		Container:      container,
	}

	container.Filter(cors.Filter)
	container.Filter(container.OPTIONSFilter)

	return container, nil
}

func main() {
	restContainer, err := Factory()
	if err != nil {
		log.Fatalf("failed to [initialize] Swagger service, %s", err.Error())
	}

	log.Printf("swagger documentation is available at [%s]\n", defaultServicePort)

	err = http.ListenAndServe(defaultServicePort, restContainer)
	if err != nil {
		log.Fatalf("failed to [serve] Swagger service, %s", err.Error())
	}
}
