package query

const (
	CreateProduct = `
	INSERT INTO products (
		ulid, 
		name, 
		price, 
		description, 
		condition, 
		updated_at, 
		deleted_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7
	) RETURNING *;
`
)