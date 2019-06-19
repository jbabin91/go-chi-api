package main

import (
	"fmt"
	"go-chi-api/httpd/handler"
	"go-chi-api/platform/newsfeed"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":3000"
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		Title: "Hello",
		Post:  "World",
	})

	r := chi.NewRouter()

	// Get newsfeed
	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	// Post newsfeed
	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on http://localhost" + port)
	http.ListenAndServe(port, r)
}
