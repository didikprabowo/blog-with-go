package models

import (
	"database/sql"
	"fmt"
	"github.com/didikprabowo/blog/database"
)

type Post struct {
	Id          int
	Title       string
	Slug        string
	Description string
	Content     string
	Image       string
	Category    string
	Created_at  string
}

func GetAllPost() *sql.Rows {
	db := database.MySQL()
	query, err := db.Query("SELECT posts.id, posts.title,posts.slug,posts.description,posts.content ," +
		"posts.image,categories.name, posts.created_at FROM posts " +
		"INNER JOIN categories  ON posts.category_id = categories.id")
	if err != nil {
		panic(err.Error)
	}
	return query
}
func DetailPost(slug string) *sql.Rows {
	db := database.MySQL()
	result := fmt.Sprintf("SELECT posts.id, posts.title,posts.description,posts.content ,"+
		"posts.image,categories.name, posts.created_at FROM posts "+
		"INNER JOIN categories  ON posts.category_id = categories.id where posts.slug = %q", slug)

	query, err := db.Query(result)
	if err != nil {
		panic(err.Error())
	}
	return query
}
