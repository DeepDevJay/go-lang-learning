package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

/*
mapping like this

	 FS32FSD23  --> {
		"ID": "23232",
		"OriginalURL": "https://www.google.com",
		"ShortURL": "http://localhost:8080/FS32FSD23",
		"CreationDate": "2021-09-01T12:00:00Z"
	 }
*/
var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // It converts the original URL to byte slice
	// fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	// fmt.Println("hasher data:", data)
	hash := hex.EncodeToString(data)
	// fmt.Println("EncodedToString data:", hash)
	fmt.Println("final string:", hash[:8])
	return hash[:8]
}

func createURL(OriginalURL string) string {
	shortUrl := generateShortURL(OriginalURL)
	id := shortUrl
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortUrl,
		CreationDate: time.Now(),
	}
	return shortUrl
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func RootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func ShourtURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortURL_ := createURL(data.URL)
	// fmt.Fprintf(w, shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	// fmt.Println("Starting URL Shortner....")
	// OriginalURL := "https://www.google.com"
	// generateShortURL(OriginalURL)

	// Register the handler function to handle all requests to the root URL ("/")
	http.HandleFunc("/", RootPageURL)
	http.HandleFunc("/shorten", ShourtURLHandler)
	http.HandleFunc("/redirect/", RedirectURLHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server started on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error while starting the server: ", err)
	}
}
