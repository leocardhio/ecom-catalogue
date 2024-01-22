package model

import "github.com/leocardhio/ecom-catalogue/datastruct"

type CreateProductRequest struct {
	Name        string                      `json:"name" binding:"required"`
	Price       uint                        `json:"price" binding:"required"`
	Tags        []datastruct.Tag            `json:"tags"`
	ImageUrls   []string                    `json:"imageUrls" binding:"required"`
	Description string                      `json:"description" binding:"required"`
	Condition   datastruct.ProductCondition `json:"condition" binding:"required"`
}

type CreateProductResponse struct {
	Id    string `json:"id" binding:"required"`
	Count int64  `json:"count" binding:"required"`
}

type GetProductRequest struct {
	Id string `uri:"id" binding:"required"`
}

type GetProductResponse struct {
	datastruct.Product
}

type GetProductsRequest struct {
	PageSize   uint `json:"limit" binding:"required"`
	PageNumber uint `json:"offset" binding:"required"`
}

type GetProductsResponse struct {
	Products  []datastruct.Product `json:"products" binding:"required"`
	Count     uint                 `json:"count" binding:"required"`
	Page      uint                 `json:"page" binding:"required"`
	TotalPage uint                 `json:"total_page" binding:"required"`
}

type UpdateProductRequestURI struct {
	Id string `uri:"id" binding:"required"`
}

type UpdateProductRequestBody struct {
	Name        string                      `json:"name" binding:"required"`
	Price       uint                        `json:"price" binding:"required"`
	Tags        []datastruct.Tag            `json:"tags"`
	ImageUrls   []string                    `json:"imageUrls" binding:"required"`
	Description string                      `json:"description" binding:"required"`
	Condition   datastruct.ProductCondition `json:"condition" binding:"required"`
}

type UpdateProductRequest struct {
	UpdateProductRequestURI
	UpdateProductRequestBody
}

type UpdateProductResponse struct {
	Count int64 `json:"count" binding:"required"`
}

type UpdateProductStatusRequestURI struct {
	Id string `uri:"id" binding:"required"`
}

type UpdateProductStatusRequestBody struct {
	IsSold bool `json:"is_sold" binding:"required"`
}

type UpdateProductStatusRequest struct {
	UpdateProductStatusRequestURI
	UpdateProductStatusRequestBody
}

type UpdateProductStatusResponse struct {
	Count int64 `json:"count" binding:"required"`
}

type DeleteProductRequest struct {
	Id string `uri:"id" binding:"required"`
}

type DeleteProductResponse struct {
	Count int64 `json:"count" binding:"required"`
}
