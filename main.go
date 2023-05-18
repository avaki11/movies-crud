package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies = []Movie{}

func getMovie(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(response).Encode(item)
			return
		}
	}

}

func deleteMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(movies)
}

func getAllMovies(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(movies)

}

func main() {
	router := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "Hello", Director: &Director{Firstname: "John", Lastname: "Pet"}})
	movies = append(movies, Movie{ID: "2", Isbn: "656", Title: "Commando", Director: &Director{Firstname: "Daniel", Lastname: "William"}})
	movies = append(movies, Movie{ID: "3", Isbn: "1789", Title: "Rush Hour 1", Director: &Director{Firstname: "Jeff", Lastname: "Elder"}})
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", updateMovieByID).Methods("PUT")
	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
