package main

import (
	"fmt"
	"go-rest-api/controllers"
	"go-rest-api/database"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux

// var db *sql.DB
var db = database.GetDb()

func routers() *chi.Mux {
	router.Get("/posts", controllers.Index)
	router.Get("/posts/{id}", controllers.Show)
	router.Post("/posts", controllers.Store)
	router.Put("/posts/{id}", controllers.Update)
	router.Delete("/posts/{id}", controllers.Delete)

	return router
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
}

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}
