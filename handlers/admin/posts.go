package admin

import (
	"fmt"
	"github.com/didikprabowo/blog/database"
	"github.com/didikprabowo/blog/models"
	"github.com/grokify/html-strip-tags-go" // => strip
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db := database.MySQL()
	query, err := db.Query("SELECT posts.id, posts.title,posts.slug,posts.description,posts.content,posts.image,categories.name FROM posts INNER JOIN categories  ON posts.category_id = categories.id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var posts []models.Post
	for query.Next() {
		var post models.Post
		var description string
		query.Scan(&post.Id, &post.Title, &post.Slug, &description, &post.Content, &post.Image, &post.Category)
		post.Description = strip.StripTags(description)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
func StorePost(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	content := r.PostFormValue("content")
	category := r.PostFormValue("category_id")
	slug := strings.ToLower(strings.Replace(title, " ", "-", -1))

	file, handler, err := r.FormFile("uploadfile")
	db := database.MySQL()
	query, err := db.Prepare("INSERT INTO posts (title,description,content,image,category_id,created_at,slug) values(?,?,?,?,?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query.Exec(title, description, content, "assets/img/"+handler.Filename, category, time.Now(), slug)
	//upload image
	f, err := os.OpenFile("assets/img/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(f, file)
	http.Redirect(w, r, "/admin/post/", 301)
}
