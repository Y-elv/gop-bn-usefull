package routes

import (
	"github.com/Y-elv/gop-bn-usefull.git/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Define article routes and their handlers
	r.HandleFunc("/articles", handlers.GetAllArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", handlers.GetArticleByID).Methods("GET")
	r.HandleFunc("/articles", handlers.CreateArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", handlers.UpdateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", handlers.DeleteArticle).Methods("DELETE")

	return r
}