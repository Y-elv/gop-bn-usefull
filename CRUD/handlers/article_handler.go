package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Y-elv/gop-bn-usefull.git/common"
	"github.com/Y-elv/gop-bn-usefull.git/models"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get all articles")
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Get article by ID")
}

// CreateArticleHandler handles POST requests to create a new article.

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	var b models.Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid body"})
		return
	}

	// Create the book
	coll := common.GetDBCollection("articles")
	result, err := coll.InsertOne(r.Context(), b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Failed to create article",
			"message": err.Error(),
		})
		return
	}

	// Return the book as JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": result,
		"data":b,
	})
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
