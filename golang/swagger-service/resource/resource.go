package resource

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"github.com/ic2hrmk/snips/golang/swagger-service/resource/representation"
	"github.com/ic2hrmk/snips/golang/swagger-service/service"
)

type BaseResource struct {
	baseService  *service.BaseService
	webContainer *restful.Container
}

func NewBaseResourceWithSwagger(
	baseService *service.BaseService,
) *BaseResource {
	baseResource := &BaseResource{
		baseService:  baseService,
		webContainer: restful.NewContainer(),
	}
	//
	// Create base web resource
	//
	baseWebService := NewBaseWebService(baseResource)

	//
	// Initialize swagger web service
	//
	c := restfulspec.Config{
		WebServices: []*restful.WebService{
			baseWebService,
		},
		WebServicesURL: "https://base.com/",
		APIPath:        "/swagger/apidocs.json",
		PostBuildSwaggerObjectHandler: func(swo *spec.Swagger) {
			swo.Info = &spec.Info{
				InfoProps: spec.InfoProps{
					Title: "Base API docs",
					Description: "Basic documentation for 'Base Service'",
					Contact: &spec.ContactInfo{
						Name: "Base Team",
					},
					Version: "v0.1.0",
				},
			}
		},
	}

	swaggerWebService := restfulspec.NewOpenAPIService(c)

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		CookiesAllowed: false,
		Container:      baseResource.webContainer,
	}

	baseResource.webContainer.Filter(cors.Filter)
	baseResource.webContainer.Filter(baseResource.webContainer.OPTIONSFilter)

	//
	// Attach web services
	//
	baseResource.webContainer.Add(baseWebService)
	baseResource.webContainer.Add(swaggerWebService)

	return baseResource
}

func NewBaseWebService(
	resource *BaseResource,
) (
	*restful.WebService,
) {
	ws := &restful.WebService{}

	tags := []string{"Base resource"}

	ws.
		Path("/bases").
		Doc("Base Resource")

	ws.Route(ws.POST("").
		To(resource.Create).
		Doc("Create base").
		Operation("base.create").
		Reads(representation.CreateBaseRequest{}).
		Writes(representation.CreateBaseResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.CreateBaseResponse{}))

	return ws
}

func (rcv *BaseResource) Create(request *restful.Request, response *restful.Response) {
	log.Println("Base creation request")
}

func (rcv *BaseResource) RunStandalone(address string) error {
	return http.ListenAndServe(address, rcv.webContainer)
}
