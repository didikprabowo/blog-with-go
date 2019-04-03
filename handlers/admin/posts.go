package admin

import (
	"github.com/didikprabowo/blog/database"
	"github.com/didikprabowo/blog/models"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db := database.MySQL()
	query, err := db.Query("SELECT posts.id, posts.title,posts.slug,posts.description,posts.content,posts.image,categories.name FROM posts INNER JOIN categories  ON posts.category_id = categories.id")
	if err != nil {
		panic("Query Error Om" + err.Error())
	}
	var posts []models.Post
	for query.Next() {
		var post models.Post
		query.Scan(&post.Id, &post.Title, &post.Slug, &post.Description, &post.Content, &post.Image, &post.Category)
		posts = append(posts, post)
	}

	Load := map[string]interface{}{
		"Results": posts,
		"Titles":  Meta{Title: "Posts"},
	}
	tmpl.ExecuteTemplate(w, "posts.html", Load)
}
func CreatePost(w http.ResponseWriter, r *http.Request) {

	db := database.MySQL()
	categoriesQ, err := db.Query("SELECT id,name,description,slug FROM categories ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	var categories []models.Category
	for categoriesQ.Next() {
		var categoryq models.Category
		categoriesQ.Scan(&categoryq.ID, &categoryq.Name, &categoryq.Description, &categoryq.Slug)
		categories = append(categories, categoryq)
	}
	Load := map[string]interface{}{
		"Results": categories,
		"Titles":  Meta{Title: "Create Post"},
	}
	tmpl.ExecuteTemplate(w, "CratePosts.html", Load)
}
