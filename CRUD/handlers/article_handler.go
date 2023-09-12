package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Y-elv/gop-bn-usefull.git/common"
	"github.com/Y-elv/gop-bn-usefull.git/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	coll := common.GetDBCollection("articles")

	articles := make([]models.Article, 0)
	cursor, err := coll.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for cursor.Next(r.Context()) {
		article := models.Article{}
		err := cursor.Decode(&article)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the URL
	id := mux.Vars(r)["id"]

	// Convert the ID to an ObjectID if you're using MongoDB
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the collection
	coll := common.GetDBCollection("articles")

	// Find the article by ID
	var article models.Article
	err = coll.FindOne(r.Context(), bson.M{"_id": objectID}).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Return the article as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
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

	// Create the article
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

	// Return the article as JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": result,
		"data":   b,
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
