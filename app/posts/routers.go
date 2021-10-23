package posts

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type PostRouter struct {
	Router   *mux.Router
	handlers *PostHandlerMux
}

func NewPostRouter(db *sql.DB) (*PostRouter, error) {
	postRouter := PostRouter{
		Router:   mux.NewRouter(),
		handlers: NewPostHandlerMux(db),
	}

	return &postRouter, nil
}
