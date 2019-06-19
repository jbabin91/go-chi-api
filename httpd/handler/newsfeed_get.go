package handler

import (
	"net/http"
	"encoding/json"
	"go-chi-api/platform/newsfeed"
)

// NewsfeedGet returns all items in the newsfeed
func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}