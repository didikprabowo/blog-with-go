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

func GetAllPost(mulai int, halaman int) (*sql.Rows, int) {

	db := database.MySQL()
	result := fmt.Sprintf("SELECT  posts.id, posts.title,posts.slug,posts.description,posts.content ,"+
		"posts.image,categories.name, posts.created_at FROM posts "+
		"INNER JOIN categories  ON posts.category_id = categories.id order By id DESC Limit %v,%v", mulai, halaman)

	query, err := db.Query(result)
	if err != nil {
		panic(err.Error)
	}
	var count int
	row := db.QueryRow("SELECT COUNT(*) as count FROM posts;")
	row.Scan(&count)
	if err != nil {
		panic(err.Error)
	}
	return query, count
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
