package main

import (
	"database/sql"
	"log"
	"net/http"
	"postapi/app/database"
	"postapi/app/helpers"
	"postapi/app/posts"

	"github.com/gorilla/mux"
)

func main() {
	// get DB
	db := database.DB{}
	err := db.Open()
	defer db.GetDB().Close()
	helpers.CheckAndExit(err)

	// routes
	router := mux.NewRouter()
	configRoutes(db.GetDB(), router)
	http.HandleFunc("/", router.ServeHTTP)

	// start server http
	log.Println("App running...")
	err = http.ListenAndServe(":9000", nil)
	helpers.CheckAndExit(err)
}

func configRoutes(db *sql.DB, router *mux.Router) {
	postRouter := posts.NewPostHandlerMux(db)
	router.HandleFunc("/api/posts", postRouter.CreatePostHandler()).Methods("POST")
	router.HandleFunc("/api/posts", postRouter.GetPostsHandler()).Methods("GET").Queries("author", "{author}")
	router.HandleFunc("/api/posts/{post_id}", postRouter.UpdatePostHandler()).Methods("PUT")
	router.HandleFunc("/api/posts/{post_id}", postRouter.DeletePostHandler()).Methods("DELETE")
}
