package query

const (
	CreateProduct = `
	INSERT INTO products (
		ulid, 
		name, 
		price, 
		description, 
		condition
	) VALUES (
		$1, $2, $3, $4, $5
	) RETURNING *
`
)