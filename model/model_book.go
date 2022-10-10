package model

type Book struct {
	Title     string `bson:"title" json:"title" example:"Test book"`
	Author    string `bson:"author" json:"author" example:"blackironj"`
	Publisher string `bson:"publisher" json:"publisher" example:"blackironj company"`
	ISBN      string `bson:"isbn" json:"isbn" example:"000-0-00-000000-0"`
}
