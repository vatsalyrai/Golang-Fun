package main

//"fmt"
// "log"
// "encoding/json"
// "math/rand"
// "net/http"
// "strcopy"
// "github/gorilla/mux"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     int       `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

var movies []Movie

func main() {

}
