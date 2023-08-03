package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func checkURL(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World!!!!"))
	keyParam := chi.URLParam(r, "key")

	//w.Write([]byte(keyParam))
	if keyParam == "hello" {
		w.Write([]byte("World!"))
	} else if keyParam == "world" {
		w.Write([]byte("Hello!"))
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the requested URL from the request
	requestedURL := r.URL.String()

	// Print the requested URL
	fmt.Println("Requested URL:", requestedURL)

	// Send a response back to the client
	fmt.Fprintf(w, "Requested URL: %s", requestedURL)

	//paramArr := [3]string{"/", "hello", "world"}
	for i := 0; i < 3; i++ {

	}

}

func main() {
	r := chi.NewRouter()
	// 1 ////////////////////////////////////////////////////////////
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!!!"))
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World!"))
	})
	r.Get("/world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	// 2 ////////////////////////////////////////////////////////////

	//http.HandleFunc("/", handler)

	//r.Get("/queryparam", checkURL)
	//http.HandleFunc("/", handler)

	//
	http.ListenAndServe(":3000", r)

	// fmt.Println("Hello, World!!")

}
