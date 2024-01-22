package query

const (
	CreateProduct = `
	INSERT INTO products (
		id, 
		name, 
		price, 
		description, 
		condition
	) VALUES (
		$1, $2, $3, $4, $5
	)
`

	CreateProductTags = `
	INSERT INTO product_tags (
		product_id,
		tag_id
	) VALUES (
		$1, $2
	)
	`

	CreateProductImage = `
	INSERT INTO product_images (
		product_id,
		image_url
	) VALUES (
		$1, $2
	)
	`

	GetProductImageUrlsByProductId = `
	SELECT image_url
	FROM product_images
	WHERE product_id = $1
	`

	DeleteProductTags = `
	DELETE FROM product_tags
	WHERE product_id = $1 AND tag_id = $2
	`

	UpdateProduct = `
	UPDATE products
	SET name = $1, price = $2, description = $3, condition = $4, updated_at = NOW()
	WHERE id = $5 AND deleted_at IS NULL
	`

	UpdateProductStatus = `
	UPDATE products
	SET is_sold = $1, updated_at = NOW()
	WHERE id = $2 AND deleted_at IS NULL
	`

	DeleteProduct = `
	UPDATE products
	SET deleted_at = NOW()
	WHERE id = $1 AND deleted_at IS NULL
	`

	DeleteProductImage = `
	DELETE FROM product_images
	WHERE product_id = $1 AND image_url = $2
	`
)
