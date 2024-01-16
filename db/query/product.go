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
	) RETURNING *
`

	UpdateProduct = `
	UPDATE products
	SET name = $1, price = $2, description = $3, condition = $4, updated_at = NOW()
	WHERE id = $5
	RETURNING *
	`
)
