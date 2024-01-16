package query

const (
	CreateTag = `
		INSERT INTO tags (
			name
		) VALUES (
			$1
		) RETURNING id
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
