package models



// Article represents the structure of an article document in the database.
type Article struct {
	ID        string             `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title" `
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author" `
	Image     string             `json:"image" bson:"image" `
	       
}


