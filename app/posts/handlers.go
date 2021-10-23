package posts

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"postapi/app/helpers"
	"strconv"
)

type PostHandlers interface {
	CreatePostHandler() http.HandlerFunc
	GetPostsHandler() http.HandlerFunc
	UpdatePostHandler() http.HandlerFunc
	DeletePostHandler() http.HandlerFunc
}

type PostHandlerMux struct {
	usecase PostUsecase
}

func NewPostHandlerMux(db *sql.DB) *PostHandlerMux {
	return &PostHandlerMux{
		usecase: NewDBPostUsecase(db),
	}
}

func (handler *PostHandlerMux) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get from body
		req := PostRequest{}
		err := helpers.Parse(w, r, &req)
		if err != nil {
			helpers.SendResponse(w, r, "", nil, http.StatusBadRequest)
			log.Printf("Cannot parse post body. err=%v \n", err)
			return
		}

		// Create the post
		post, err := NewPost(Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		})
		if err != nil {
			helpers.SendResponse(w, r, err.Error(), nil, http.StatusBadRequest)
			return
		}

		// create post
		post, err = handler.usecase.Create(post)

		// return post json
		postJson := post.ToJSON()
		helpers.SendResponse(w, r, "", postJson, http.StatusCreated)
	}
}

func (handler *PostHandlerMux) GetPostsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get author from query param
		vars := helpers.GetVarsRoute(r)
		author, found := vars["author"]
		if found == false {
			msg := fmt.Sprint("missing post_id from update post.")
			helpers.SendResponse(w, r, msg, nil, http.StatusBadRequest)
			return
		}

		// get posts
		posts, err := handler.usecase.GetAll(author)
		if err != nil {
			msg := fmt.Sprint("Cannot get posts.")
			helpers.SendResponse(w, r, msg, nil, http.StatusInternalServerError)
			return
		}

		// to json
		var postsJSON = make([]*JsonPost, len(posts))
		for index, post := range posts {
			postsJSON[index] = post.ToJSON()
		}
		helpers.SendResponse(w, r, "", postsJSON, http.StatusOK)
	}
}

func (handler *PostHandlerMux) UpdatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := helpers.GetVarsRoute(r)
		postIDStr, found := vars["post_id"]
		if found == false {
			msg := fmt.Sprint("missing post_id from update post.")
			helpers.SendResponse(w, r, msg, nil, http.StatusBadRequest)
			return
		}
		postID, _ := strconv.Atoi(postIDStr)

		// get post from body
		req := PostRequest{}
		err := helpers.Parse(w, r, &req)
		if err != nil {
			helpers.SendResponse(w, r, err.Error(), nil, http.StatusBadRequest)
			return
		}

		// fetch post
		postFound, err := handler.usecase.Get(postID)
		if err != nil {
			helpers.SendResponse(w, r, err.Error(), nil, http.StatusInternalServerError)
			return
		}
		if postFound == nil {
			helpers.SendResponse(w, r, errors.New("post not found").Error(), nil, http.StatusInternalServerError)
			return
		}

		// create post model
		post, err := NewPost(Post{
			ID:      postID,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		})
		if err != nil {
			msg := fmt.Sprint("Cannot parse post body.")
			helpers.SendResponse(w, r, msg, nil, http.StatusBadRequest)
			return
		}

		// Update post
		err = handler.usecase.Update(post)

		// return no content
		helpers.SendResponse(w, r, "", nil, http.StatusNoContent)
	}
}

func (handler *PostHandlerMux) DeletePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get id from params
		vars := helpers.GetVarsRoute(r)
		postIDStr, found := vars["post_id"]
		if found == false {
			msg := fmt.Sprint("Cannot parse post body.")
			helpers.SendResponse(w, r, msg, nil, http.StatusBadRequest)
			return
		}
		postID, _ := strconv.Atoi(postIDStr)

		// fetch post
		postFound, err := handler.usecase.Get(postID)
		if err != nil {
			helpers.SendResponse(w, r, err.Error(), nil, http.StatusInternalServerError)
			return
		}
		if postFound == nil {
			helpers.SendResponse(w, r, errors.New("post not found").Error(), nil, http.StatusInternalServerError)
			return
		}

		// delete
		err = handler.usecase.Delete(postID)
		if err != nil {
			msg := fmt.Sprintf("Cannot delete post in DB. err=%v", err)
			helpers.SendResponse(w, r, msg, nil, http.StatusBadRequest)
			return
		}

		// return no content
		helpers.SendResponse(w, r, "", nil, http.StatusNoContent)
	}
}
