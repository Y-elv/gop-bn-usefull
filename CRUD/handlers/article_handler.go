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
	fmt.Println(id)

	// Convert the ID to an ObjectID if you're using MongoDB
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the collection
	coll := common.GetDBCollection("articles")

	// Find the article by ID
	var article models.Article
	err = coll.FindOne(r.Context(), bson.M{"_id": id}).Decode(&article)
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
		"comment":"article created successfully",

	})
	
}
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
    // Get the ID parameter from the URL
    id := mux.Vars(r)["id"]

    // Convert the ID to an ObjectID if you're using MongoDB
    _, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get the collection
    coll := common.GetDBCollection("articles")

    // Define the update filter (in this example, we're using the article's ID)
    filter := bson.M{"_id": id}

    // Parse the request body to get the update data
    var updateData map[string]interface{}
    err = json.NewDecoder(r.Body).Decode(&updateData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create the update query based on the parsed data
    updateQuery := bson.M{
        "$set": updateData,
    }

    // Perform the update operation
    updateResult, err := coll.UpdateOne(r.Context(), filter, updateQuery)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Check if any documents were modified
    if updateResult.ModifiedCount == 0 {
        http.Error(w, "No article found with the provided ID", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Article updated successfully"))
}


func DeleteArticle(w http.ResponseWriter, r *http.Request) {
    // Get the ID parameter from the URL
    id := mux.Vars(r)["id"]

    // Convert the ID to an ObjectID if you're using MongoDB
   _, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get the collection
    coll := common.GetDBCollection("articles")

    // Define the deletion filter
    filter := bson.M{"_id": id}

    // Perform the delete operation
    deleteResult, err := coll.DeleteOne(r.Context(), filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Check if any documents were deleted
    if deleteResult.DeletedCount == 0 {
        http.Error(w, "No article found with the provided ID", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Article deleted successfully"))
}

func DeleteAllArticle(w http.ResponseWriter, r *http.Request) {
   // Get the collection
   coll := common.GetDBCollection("articles")

   // Define an empty filter to match all documents
   filter := bson.M{}

   // Perform the delete operation
   deleteResult, err := coll.DeleteMany(r.Context(), filter)
   if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
   }

   // Check if any documents were deleted
   if deleteResult.DeletedCount == 0 {
	   http.Error(w, "No articles found to delete", http.StatusNotFound)
	   return
   }

   w.WriteHeader(http.StatusOK)
   w.Write([]byte("All articles deleted successfully"))
}
