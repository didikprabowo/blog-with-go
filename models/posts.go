package models

import (
	"database/sql"
	"fmt"
	"github.com/didikprabowo/blog/database"
)

type Post struct {
	Id           int
	Title        string
	Slug         string
	Description  string
	Content      string
	Image        string
	Category     string
	Created_at   string
	SlugCategory string
}

var db = database.MySQL()

func GetAllPost(mulai int, halaman int, slugCat string) (*sql.Rows, int) {
	var slug string
	slug = "%" + slugCat + "%"
	result := fmt.Sprintf("SELECT  posts.id, posts.title,posts.slug,posts.description,posts.content ,"+
		"posts.image,categories.name, posts.created_at, categories.slug FROM posts "+
		"INNER JOIN categories  ON posts.category_id = categories.id where categories.slug LIKE %q order By id DESC Limit %v,%v", slug, mulai, halaman)
	query, err := db.Query(result)
	if err != nil {
		panic(err.Error)
	}
	var count int
	resultCount := fmt.Sprintf("SELECT COUNT(*) as count FROM posts INNER JOIN categories  ON posts.category_id = categories.id where categories.slug LIKE %q", "%"+slug+"%")
	row := db.QueryRow(resultCount)
	row.Scan(&count)
	if err != nil {
		panic(err.Error)
	}
	return query, count
}
func DetailPost(slug string) *sql.Rows {
	result := fmt.Sprintf("SELECT posts.id, posts.title,posts.description,posts.content ,"+
		"posts.image,categories.name, posts.created_at FROM posts "+
		"INNER JOIN categories  ON posts.category_id = categories.id where posts.slug = %q", slug)

	query, err := db.Query(result)
	if err != nil {
		panic(err.Error())
	}
	return query
}
