package main

import (

	//"fmt"
	// "log"
	// "encoding/json"
	// "math/rand"
	// "net/http"
	// "strcopy"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string       `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)


}





func getMovie(w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	 
	for  _, item := range movies{ // we have used blank identifier because we don't need index here

		if(item.ID==params["id"])
		{
          
           
		   json.NewEncoder(w).Encode(item)
		   return
		}
	}

}


func updateMovie(w http.ResponseWriter , r *http.Request){
	


	
}

func createMovie(w http.ResponseWriter , r *http.Request){
  
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.body).Decode(movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)


}



func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	 
	for index , item := range movies{

		if(item.ID==params["id"])
		{
           movies = append(movies[:index],movies[index+1:]...)

		   break;
		}
	}

	json.NewEncoder(w).Encode(movies)
}




func main() {

	r:=mux.NewRouter()

	movies = append(movies, Movie{ID:"1",Isbn:"47635",Title:"SomeMovie1",Director : &Director{ Firstname:"First1" , Lastname:"Last1"}})
	movies = append(movies, Movie{ID:"2",Isbn:"40635",Title:"SomeMovie2",Director : &Director{ Firstname:"First2" , Lastname:"Last2"}})
	movies = append(movies, Movie{ID:"3",Isbn:"41635",Title:"SomeMovie3",Director : &Director{ Firstname:"First3" , Lastname:"Last3"}})
	movies = append(movies, Movie{ID:"4",Isbn:"49635",Title:"SomeMovie4",Director : &Director{ Firstname:"First4" , Lastname:"Last4"}})


	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("movies",createMovie).Methods("POST")
	r.HandleFunc("movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}",deleteMovie).Methods("DELETE")


	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r)) // to start the server


}
