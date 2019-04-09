package admin

import (
	"fmt"
	"github.com/didikprabowo/blog/database"
	"github.com/didikprabowo/blog/models"
	"github.com/gorilla/mux"
	"github.com/grokify/html-strip-tags-go" // => strip
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db := database.MySQL()
	query, err := db.Query("SELECT posts.id, posts.title,posts.slug,posts.description,posts.content," +
		"posts.image,categories.name FROM posts " +
		"INNER JOIN categories  ON posts.category_id = categories.id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var posts []models.Post
	for query.Next() {
		var post models.Post
		var description string
		query.Scan(&post.Id, &post.Title, &post.Slug,
			&description, &post.Content, &post.Image, &post.Category)
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
		categoriesQ.Scan(&categoryq.ID, &categoryq.Name,
			&categoryq.Description, &categoryq.Slug)
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
	query, err := db.Prepare("INSERT INTO posts (title,description,content,image,category_id,created_at,slug)" +
		"values(?,?,?,?,?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query.Exec(title, description, content, "static/img/"+handler.Filename, category, time.Now(), slug)
	//upload image
	f, err := os.OpenFile("assets/img/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(f, file)
	http.Redirect(w, r, "/admin/post/", 301)
}
func EditPost(w http.ResponseWriter, r *http.Request) {
	db := database.MySQL()
	VarID := mux.Vars(r)
	id := VarID["id"]

	result := fmt.Sprintf("SELECT id, title,description,content,image,category_id FROM posts where id = %v", id)

	query, err := db.Query(result)
	if err != nil {
		panic(err.Error())
	}
	post := models.Post{}
	for query.Next() {
		query.Scan(&post.Id, &post.Title, &post.Description, &post.Content, &post.Image, &post.Category)
	}
	//category
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

	catid, _ := strconv.Atoi(post.Category)
	Load := map[string]interface{}{
		"Catid":    catid,
		"Results":  post,
		"Category": categories,
		"Titles":   Meta{Title: "Edit Post"},
	}
	tmpl.ExecuteTemplate(w, "EditPosts.html", Load)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")

	var title, description, content, category_id, image string

	title = r.PostFormValue("title")
	description = r.PostFormValue("description")
	content = r.PostFormValue("content")
	category_id = r.PostFormValue("category_id")
	file, handler, err := r.FormFile("uploadfile")

	if err != nil {
		image = r.PostFormValue("old_image")
	} else {
		f, _ := os.OpenFile("assets/img/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		image = "static/img/" + handler.Filename
		io.Copy(f, file)
		old := strings.Replace(r.PostFormValue("old_image"), "static", "assets", -1)
		os.Remove(old)
	}
	db := database.MySQL()
	result := fmt.Sprintf("UPDATE posts set image = %q, title = %q , description = %q, content = %q , category_id = %q where id = %v", image, title, description, content, category_id, id)
	db.Query(result)

	http.Redirect(w, r, "/admin/post", 301)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := database.MySQL()
	VarID := mux.Vars(r)
	id := VarID["id"]

	result := fmt.Sprintf("DELETE from posts where id = %v", id)
	db.Query(result)
	fmt.Println(result)
	http.Redirect(w, r, "/admin/post", 301)
}
