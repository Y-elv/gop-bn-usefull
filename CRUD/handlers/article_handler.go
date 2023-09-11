package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	// Parse the JSON request body into an Article model
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simulate the comment parameter from the request URL (you can handle this differently)
	comment := r.URL.Query().Get("comment")

	// Set additional fields for the article
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	// Save the article to the database
	title := "Sample Title"
	content := "Sample Content"
	author := "Sample Author"
	image := "Sample Image URL"
	if err := models.NewArticle(title, content, author, image); err != nil {
		http.Error(w, "Failed to save the data", http.StatusInternalServerError)
		return
	}

	// Respond with a JSON success message
	response := struct {
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
		Data    interface{} `json:"data"`
		Comment string      `json:"comment"`
	}{
		Message: "Data saved successfully",
		Error:   nil,
		Data:    article,
		Comment: comment,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
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
