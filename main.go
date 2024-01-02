package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	m := http.NewServeMux()
	const addr = ":80"
	m.HandleFunc("/", handler)

	srv := http.Server{
		Addr:         addr,
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// Return an HTML page with a random paragraph
	page := "<html><body><p>Your random string: " + RandomStr(10) + "</p></body></html>"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page))
}

/**
 * Generate a random string of a given length
 */
func RandomStr(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
