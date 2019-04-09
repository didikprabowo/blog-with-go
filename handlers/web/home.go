package web

import (
	"github.com/didikprabowo/blog/database"
	"github.com/didikprabowo/blog/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Beranda(w http.ResponseWriter, r *http.Request) {

	halamanFormUri := r.URL.Query().Get("halaman")
	var halaman int
	halaman = 10
	var page, mulai int
	if len(halamanFormUri) < 1 {
		mulai = 0
	} else {
		ke, _ := strconv.Atoi(halamanFormUri)
		if ke == 0 {
			mulai = 0
		} else {
			page = ke
			mulai = page*halaman - halaman
		}
	}

	var posts []models.Post
	query, count := models.GetAllPost(mulai, halaman, "")
	var post models.Post
	for query.Next() {

		var description string
		query.Scan(&post.Id, &post.Title, &post.Slug,
			&description, &post.Content, &post.Image,
			&post.Category, &post.Created_at, &post.SlugCategory)
		post.Description = description[0:50]
		posts = append(posts, post)
	}
	var paging int
	paging = count / halaman

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
		"Results":     posts,
		"Title":       "Beranda",
		"Pagings":     paging,
		"CategoryAll": categories,
	}
	tmpl.ExecuteTemplate(w, "beranda.html", Load)
}

// Post By categories
func PostByCategory(w http.ResponseWriter, r *http.Request) {
	VarID := mux.Vars(r)
	slugCat := VarID["slug"]
	halamanFormUri := r.URL.Query().Get("halaman")
	var halaman int
	halaman = 10
	var page, mulai int
	if len(halamanFormUri) < 1 {
		mulai = 0
	} else {
		ke, _ := strconv.Atoi(halamanFormUri)
		if ke == 0 {
			mulai = 0
		} else {
			page = ke
			mulai = page*halaman - halaman
		}
	}

	var posts []models.Post
	query, count := models.GetAllPost(mulai, halaman, slugCat)
	var post models.Post
	for query.Next() {

		var description string
		query.Scan(&post.Id, &post.Title, &post.Slug,
			&description, &post.Content, &post.Image,
			&post.Category, &post.Created_at, &post.SlugCategory)
		post.Description = description[0:50]
		posts = append(posts, post)
	}
	var paging int
	paging = count / halaman

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
		"Results":     posts,
		"Title":       post.Category,
		"Pagings":     paging,
		"CategoryAll": categories,
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
