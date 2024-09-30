package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

// mapping short url to URL struct
type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"orignial_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:6]
}
func saveURL(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	urlDB[shortURL] = URL{
		ID:           shortURL,
		OriginalURL:  OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(shortURL string) (URL, error) {
	url, ok := urlDB[shortURL]
	if !ok {
		return URL{}, fmt.Errorf("URL not found")
	}
	return url, nil
}
func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		OriginalURL := r.Form.Get("url")
		shortURL := saveURL(OriginalURL)

		fmt.Fprintf(w, "Short URL: %s", shortURL)

	}
}
func main() {
	fmt.Print("URL Shortening service starting\n")
	http.HandleFunc("/", handler)
	fmt.Print("Server started at port 8080\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Error starting server %v", err)
	}
}
