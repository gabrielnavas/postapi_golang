package posts

import (
	"database/sql"
	"log"
)

type PostRepository interface {
	CreatePost(p *Post) error
	UpdatePost(post *Post) error
	DeletePost(postID int) error
	Get(postID int) (*Post, error)
	GetPostsByAuthors(authors string) ([]*Post, error)
	GetPosts() ([]*Post, error)
}

type PostRepositoryDB struct {
	db *sql.DB
}

func NewPostRepositoryDB(db *sql.DB) *PostRepositoryDB {
	return &PostRepositoryDB{
		db,
	}
}

func (p *PostRepositoryDB) CreatePost(post *Post) error {
	row := p.db.QueryRow(InsertPostSchema, post.Title, post.Content, post.Author)
	err := row.Scan(&post.ID)
	if err != nil {
		log.Printf("Cannot insert post. err=%v", err)
		return err
	}
	return nil
}

func (p *PostRepositoryDB) GetPosts() ([]*Post, error) {
	var posts []*Post
	var rows *sql.Rows

	rows, err := p.db.Query(SelectPostsSchema)
	defer rows.Close()
	log.Printf("Cannot query posts. err=%v", err)
	if err != nil {
		return posts, err
	}

	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *PostRepositoryDB) GetPostsByAuthors(author string) ([]*Post, error) {
	var posts []*Post
	var rows *sql.Rows

	stmt, err := p.db.Prepare(SelectPostsSchemaWithLike)
	defer stmt.Close()
	if err != nil {
		log.Printf("Cannot prepare posts. err=%v", err)
		return posts, err
	}
	author = "%" + author + "%"
	rows, err = stmt.Query(author)
	defer rows.Close()
	if err != nil {
		log.Printf("Cannot query posts with %v. err=%v", author, err)
		return posts, err
	}

	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *PostRepositoryDB) UpdatePost(post *Post) error {
	_, err := p.db.Exec(UpdatePostSchema, post.ID, post.Title, post.Content, post.Author)
	if err != nil {
		log.Printf("Cannot update post. err=%v", err)
		return err
	}
	return nil
}

func (p *PostRepositoryDB) DeletePost(postID int) error {
	_, err := p.db.Exec(DeletePostSchema, postID)
	if err != nil {
		log.Printf("Cannot delete post. err=%v", err)
		return err
	}
	return nil
}

func (p *PostRepositoryDB) Get(postID int) (*Post, error) {
	row := p.db.QueryRow(SelectOnePostSchema, postID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
	if err != nil {
		return nil, nil
	}

	return &post, nil
}
