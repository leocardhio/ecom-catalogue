package model

import (
	"time"

	"github.com/leocardhio/ecom-catalogue/class"
)

type ProductModel struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Condition   class.ProductCondition `json:"condition"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Condition   class.ProductCondition `json:"condition"`
}

type CreateProductResponse struct {
	Ulid string `json:"ulid"`
}

type GetProductRequest struct {
	Ulid string `json:"ulid"`
}

type GetProductsRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetProductsResponse struct {
	Products []ProductModel `json:"products"`
	Count    int            `json:"count"`
	Page 		int            `json:"page"`
}

type UpdateProductRequest struct {
	Ulid        string `json:"ulid"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Condition 	class.ProductCondition `json:"condition"`
}

type DeleteProductRequest struct {
	Ulid string `json:"ulid"`
}