package db

import (
	"api-postgres/internal/comment"
	"context"
	"database/sql"
	"fmt"
)

// model here
// why have exact existing in comment package?
// db schema allows insert fields that should be potentially null
// defining specific raw struct - handle null string cases
// then convert back to struct that business layer can understand
// that approach implements loose coupling btw data and bus layers
type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

// helper
func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		// map fields
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

// call method, defined in bus-layer internal/comment/comment.go
// and return to original caller
func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {

	// retrieve comment from db
	var cmtRaw CommentRow
	raw := d.Client.QueryRowContext(
		ctx,
		`SELECT id,slug,body,author
		FROM comments
		WHERE id=$1`,
		uuid,
	)

	err := raw.Scan(&cmtRaw.ID, &cmtRaw.Slug, &cmtRaw.Body, &cmtRaw.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching comment by uuid: %w", err)
	}

	// convert comment raw into comment struct

	return convertCommentRowToComment(cmtRaw), nil
}
