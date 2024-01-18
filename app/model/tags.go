package model

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateTagResponse struct {
	Id    string `json:"id" binding:"required"`
	Count int64  `json:"count" binding:"required"`
}

type UpdateTagRequestURI struct {
	Id string `uri:"id" binding:"required"`
}

type UpdateTagRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagRequest struct {
	UpdateTagRequestURI
	UpdateTagRequestBody
}

type UpdateTagResponse struct {
	Count int64 `json:"count" binding:"required"`
}

type DeleteTagRequest struct {
	Id string `uri:"id" binding:"required"`
}

type DeleteTagResponse struct {
	Count int64 `json:"count" binding:"required"`
}
