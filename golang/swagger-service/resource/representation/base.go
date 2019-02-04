package representation

type CreateBaseRequest struct {
	Dimension   string `json:"dimension" description:"Base dimension"`
	Description string `json:"description,omitempty" description:"Base description"`
}

type CreateBaseResponse struct {
	ID          int    `json:"id" description:"Unique object identifier" type:"integer"`
	Dimension   string `json:"dimension" description:"Base dimension"`
	Description string `json:"description,omitempty" description:"Base description"`
}
