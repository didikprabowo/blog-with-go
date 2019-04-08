package web

import (
	"github.com/didikprabowo/blog/database"
	"github.com/didikprabowo/blog/models"
	"github.com/gorilla/mux"
	"github.com/grokify/html-strip-tags-go"
	"net/http"
	"strconv"
)

func Beranda(w http.ResponseWriter, r *http.Request) {

	var posts []models.Post
	query := models.GetAllPost()
	for query.Next() {
		var post models.Post
		var description string
		query.Scan(&post.Id, &post.Title, &post.Slug,
			&description, &post.Content, &post.Image,
			&post.Category, &post.Created_at)
		post.Description = strip.StripTags(description[0:100])
		posts = append(posts, post)
	}

	Load := map[string]interface{}{
		"Results": posts,
		"Titles":  "Beranda",
	}
	tmpl.ExecuteTemplate(w, "beranda.html", Load)
}

func DetailPosts(w http.ResponseWriter, r *http.Request) {

	db := database.MySQL()
	VarID := mux.Vars(r)
	slug := VarID["slug"]

	query := models.DetailPost(slug)
	post := models.Post{}
	for query.Next() {
		var Content string
		query.Scan(&post.Id, &post.Title, &post.Description,
			&Content, &post.Image, &post.Category,
			&post.Created_at)
		post.Content = Content
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
		categoriesQ.Scan(&categoryq.ID, &categoryq.Name,
			&categoryq.Description, &categoryq.Slug)
		categories = append(categories, categoryq)
	}

	catid, _ := strconv.Atoi(post.Category)
	Load := map[string]interface{}{
		"Catid":    catid,
		"Results":  post,
		"Category": categories,
	}
	tmpl.ExecuteTemplate(w, "detailposts.html", Load)
}
