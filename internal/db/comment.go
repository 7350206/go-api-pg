package db

import (
	"api-postgres/internal/comment"
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
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
// !! must have twins in Service layer
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

func (d *Database) PostComment(
	ctx context.Context,
	cmt comment.Comment,
) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String() //create new uuid
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, author, body)
		VALUES
		(:id, :slug,:author, :body)`,
		postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil

}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments WHWRW id=$1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}
	return nil
}

func (d *Database) UpdateComment(
	ctx context.Context,
	id string,
	cmt comment.Comment,
) (comment.Comment, error) {

	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug =:slug,
		author=:author,
		body=:body
		WHERE id=:id`,
		cmtRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed update comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed close rows: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil

}
