package routes

import (
	"github.com/Y-elv/gop-bn-usefull.git/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Define article routes and their handlers
	r.HandleFunc("/api/v1/articles/getAll", handlers.GetAllArticles).Methods("GET")
	r.HandleFunc("/api/v1/articles/{id}", handlers.GetArticleByID).Methods("GET")
	r.HandleFunc("/api/v1/articles", handlers.CreateArticle).Methods("POST")
	r.HandleFunc("/api/v1/articles/{id}", handlers.UpdateArticle).Methods("PUT")
	r.HandleFunc("/api/v1/articles/{id}", handlers.DeleteArticle).Methods("DELETE")
	r.HandleFunc("/api/v1/articles/delete/deleteAll", handlers.DeleteAllArticle).Methods("DELETE")

	return r
}
