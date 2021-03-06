package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"jwt-ex-go/auth"
	"log"
	"net/http"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/public", public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)

	log.Println(http.ListenAndServe(":8080", r))
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "ブログのタイトル",
		Tag:   "public",
		URL:   "https://qiita.com",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "非公開ブログ",
		Tag:   "private",
		URL:   "https://qiita.com",
	}
	json.NewEncoder(w).Encode(post)
})
