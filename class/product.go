package class

import (
	"time"
)

type ProductCondition string

const (
	BRAND_NEW    ProductCondition = "Brand New"
	LIKE_NEW     ProductCondition = "Like New"
	LIKELY_USED  ProductCondition = "Likely Used"
	WELL_USED    ProductCondition = "Well Used"
	HEAVILY_USED ProductCondition = "Heavily Used"
)

type Product struct {
	Id          string
	Name        string
	Price       int
	Description string
	Condition   ProductCondition
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
