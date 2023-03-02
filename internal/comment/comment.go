package comment

import (
	"context"
	"errors"
	"fmt"
)

// - good pract is to make a comment to ANY content of package exported
// some form a comment representation
type Comment struct {
	ID     string
	Slug   string // path comment is associated with
	Body   string // comment content
	Author string
}

// define custom errors to return, nice + secure + not expose too much
var (
	ErrFetchingComments = errors.New("failed to fetch comment by id")
	ErrNotImplemented   = errors.New("not implemented yet")
)

// define an interface Repo level to implement
// and then define it as field of Service struct
// define all of methods Service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// thing to be interact through
// is the struct on which all logic will be build on top of
// create a number of methods *receiver to Service struct: get, read ...
type Service struct {
	Store Store
}

// ** constructor and composite literal patt
// returns a pointer to new service
func NewService(store Store) *Service {
	// instantiate a new service and pass a path
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("getting a comment")

	// context explain, (ctx, key, value)
	// ctx = context.WithValue(ctx, "request_id", "unique-string")
	// fmt.Println(ctx.Value("request_id")) // unique-string

	// reach out to repo layer and retrieve a comment from db
	// + dont need to know how it'll be done, just need to inplem Store interface
	// + easy unit testing by simply mock Store interface without actually talk to db
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err) // have original error here, can figure out smth about

		// - that err handling approach, when returned WHOLE err to calling user -
		// will expose too much db inplementation details,
		// return Comment{}, err

		// + make nicer
		return Comment{}, ErrFetchingComments

	}
	return cmt, nil // return comment TO WHATEVER GETS CALLING that metod
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
