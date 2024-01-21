package query

const (
	CreateTag = `
		INSERT INTO tags (
			name
		) VALUES (
			$1
		) RETURNING id
	`

	GetTags = `
		SELECT id, name
		FROM tags
		WHERE deleted_at IS NULL
	`

	GetTagsByProductId = `
		SELECT tag_id
		FROM product_tags
		WHERE product_id = $1
	`

	GetTag = `
		SELECT id, name
		FROM tags
		WHERE id = $1 AND deleted_at IS NULL
	`

	UpdateTag = `
		UPDATE tags
		SET name = $1
		WHERE id = $2 AND deleted_at IS NULL
	`

	DeleteTag = `
		UPDATE tags
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
)
