package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie

func getAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item:= range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovies(w http.ResponseWriter, r *http.Request){
	fmt.Println("getMovies")
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request){
	//set json content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)
	
	//loop over the movies, range
	// delete the move with the i.d that you've sent
	// add a new movie - the movie that we send in the body of postman

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn:"234", Title:"Movie One", Director: &Director{FirstName: "john", LastName:"Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn:"234", Title:"Movie One", Director: &Director{FirstName: "john", LastName:"Doe"}})

	fmt.Println(movies)

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}