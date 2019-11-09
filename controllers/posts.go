package controllers

import (
	"encoding/json"
	"go-rest-api/database"
	"go-rest-api/helpers"
	"net/http"

	"github.com/go-chi/chi"
)

var db = database.GetDb()

type Post struct {
	ID      int    `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

// Index gets list of posts
func Index(w http.ResponseWriter, r *http.Request) {
	post := Post{}
	posts := []Post{}

	rows, err := db.Query("select * from posts")
	helpers.Catch(err)

	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		helpers.Catch(err)
		posts = append(posts, post)
	}

	rows.Close()

	helpers.RespondwithJSON(w, http.StatusOK, posts)
}

// Show shows post
func Show(w http.ResponseWriter, r *http.Request) {

	var post Post
	id := chi.URLParam(r, "id")

	rows, err := db.Query("select * from posts where id=?", id)
	helpers.Catch(err)

	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		helpers.Catch(err)
	}

	rows.Close()

	helpers.RespondwithJSON(w, http.StatusOK, post)
}

// Store creates a new post
func Store(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Insert posts SET title=?, content=?")
	helpers.Catch(err)

	_, er := query.Exec(post.Title, post.Content)
	helpers.Catch(er)
	defer query.Close()

	helpers.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// Update updates a spesific post
func Update(w http.ResponseWriter, r *http.Request) {
	var post Post
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Update posts set title=?, content=? where id=?")
	helpers.Catch(err)
	_, er := query.Exec(post.Title, post.Content, id)
	helpers.Catch(er)

	defer query.Close()

	helpers.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})
}

// Delete removes a spesific post
func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete from posts where id=?")
	helpers.Catch(err)
	_, er := query.Exec(id)
	helpers.Catch(er)
	defer query.Close()

	helpers.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}
