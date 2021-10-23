package posts

import "errors"

type Post struct {
	ID      int
	Title   string
	Content string
	Author  string
}

var (
	ErrTitleWrong   = errors.New("title is wrong")
	ErrContentWrong = errors.New("content is wrong")
	ErrAuthorWrong  = errors.New("author is wrong")
)

func NewPost(post Post) (*Post, error) {
	if post.Title == "" {
		return nil, ErrTitleWrong
	}
	if post.Content == "" {
		return nil, ErrContentWrong
	}
	if post.Author == "" {
		return nil, ErrAuthorWrong
	}
	return &post, nil
}

func (p *Post) ToJSON() *JsonPost {
	return &JsonPost{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}

type JsonPost struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
