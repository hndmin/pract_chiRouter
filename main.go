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

func handleQueryParam(w http.ResponseWriter, r *http.Request) {
	// Get a specific query parameter by name
	paramValue := chi.URLParam(r, "paramName")

	// Alternatively, you can use r.URL.Query() to get all query parameters as a map
	queryParams := r.URL.Query()

	// Print the specific query parameter value
	fmt.Println("Query Parameter Value:", paramValue)

	// Print all query parameters
	fmt.Println("All Query Parameters:", queryParams)

	// Send a response back to the client
	fmt.Fprintf(w, "Query Parameter Value: %s\n", paramValue)
	fmt.Fprintf(w, "All Query Parameters: %+v\n", queryParams)
}

func main() {
	r := chi.NewRouter()
	// 1 ////////////////////////////////////////////////////////////
	r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!!!!"))
	// })
	// r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("World!"))
	// })
	// r.Get("/world", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello!"))
	// })
	// 2 ////////////////////////////////////////////////////////////

	//http.HandleFunc("/", handler)

	//r.Get("/queryparam", checkURL)
	//http.HandleFunc("/", handler)

	r.Get("/queryparam", handleQueryParam)

	//
	http.ListenAndServe(":3000", r)

	// fmt.Println("Hello, World!!")

}
