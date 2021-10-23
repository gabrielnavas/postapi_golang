package posts

const CreateSchema = `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title TEXT,
		content TEXT,
		author TEXT
	)
`

const InsertPostSchema = `
	INSERT INTO posts(title, content, author) 
	VALUES($1, $2, $3) 
	RETURNING id
`

const SelectPostsSchemaWithLike = `
	SELECT id, title, content, author 
	FROM posts 
	WHERE author LIKE $1
`

const SelectPostsSchema = `
	SELECT id, title, content, author 
	FROM posts
`

const SelectOnePostSchema = `
	SELECT id, title, content, author 
	FROM posts 
	WHERE id = $1
`

const UpdatePostSchema = `
	UPDATE posts
	SET title=$2, 
		content=$3, 
		author=$4
	WHERE id = $1
`

const DeletePostSchema = `
	DELETE FROM posts
	WHERE id = $1
`
