package service

import "github.com/ic2hrmk/snips/golang/swagger-service/resource/representation"

type BaseService struct {
}

func NewBaseService() *BaseService {
	return &BaseService{}
}

func (rcv *BaseService) Create(
	req *representation.CreateBaseRequest,
) (
	*representation.CreateBaseResponse, error,
) {
	return &representation.CreateBaseResponse{
		Dimension: req.Dimension,
	}, nil
}
