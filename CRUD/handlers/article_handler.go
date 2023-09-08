package handlers
import (
	"fmt"
	"net/http")

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get all articles")
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "Get article by ID")
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "Create article")
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "Update article")
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "Delete article")
}
func DeleteAllArticle(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "Delete all article")
}
