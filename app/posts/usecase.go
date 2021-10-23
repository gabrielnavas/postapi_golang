package posts

import (
	"database/sql"
	"errors"
)

type PostUsecase interface {
	Create(post *Post) (*Post, error)
	Update(post *Post) error
	Delete(postID int) error
	Get(postID int) (*Post, error)
	GetAll(author string) ([]*Post, error)
}

type DBPostUsecase struct {
	repository PostRepository
}

var (
	ErrPostWrong   = errors.New("post is wrong")
	ErrPostIDWrong = errors.New("post ID is wrong")
	ErrQueryWrong  = errors.New("query wrong")
)

func NewDBPostUsecase(db *sql.DB) *DBPostUsecase {
	return &DBPostUsecase{
		repository: NewPostRepositoryDB(db),
	}
}

func (usecase *DBPostUsecase) Create(post *Post) (*Post, error) {
	usecase.repository.CreatePost(post)
	return post, nil
}

func (usecase *DBPostUsecase) Update(post *Post) error {
	err := usecase.repository.UpdatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *DBPostUsecase) Delete(postID int) error {
	if postID <= 0 {
		return ErrPostIDWrong
	}
	return usecase.repository.DeletePost(postID)
}

func (usecase *DBPostUsecase) Get(postID int) (*Post, error) {
	if postID <= 0 {
		return nil, ErrPostIDWrong
	}
	post, err := usecase.repository.Get(postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (usecase *DBPostUsecase) GetAll(author string) ([]*Post, error) {
	var posts []*Post

	// author query
	if len(author) > 0 {
		if len(author) < 2 {
			return posts, ErrQueryWrong
		}

		posts, err := usecase.repository.GetPostsByAuthors(author)
		if err != nil {
			return posts, err
		}
		return posts, nil
	}

	// all
	posts, err := usecase.repository.GetPosts()
	if err != nil {
		return posts, err
	}
	return posts, nil
}
