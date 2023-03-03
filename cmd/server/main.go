// it'll be run first
// here will be instantiation for all components

package main

import (
	"api-postgres/internal/comment"
	"api-postgres/internal/db"
	"context"
	"fmt"
)

// responsible for instantiation and app startup
func Run() error {
	fmt.Println("connecting db...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect db")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}

	fmt.Println("db connected...")

	// test
	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
	))

	return nil
}

func main() {

	// handle errs while app starts, instead main() panic - catch err and [print]
	fmt.Println("main started")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
