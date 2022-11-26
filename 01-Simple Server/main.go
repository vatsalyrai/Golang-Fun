package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(responsew http.ResponseWriter, requestr *http.Request) {
	err := requestr.ParseForm()

	if err != nil {
		fmt.Fprint(responsew, "ParseForm() error %v ", err)
	}

	fmt.Fprintf(responsew, "POST request successful \n")

	name := requestr.FormValue("name")
	address := requestr.FormValue("address")

	fmt.Fprintf(responsew, "name is %v \n ", name)
	fmt.Fprintf(responsew, "address is %v\n", address)

}

func helloHandler(responsew http.ResponseWriter, requestr *http.Request) {

	if requestr.Method != "GET" {
		http.Error(responsew, "Method not supported", http.StatusNotFound)
	}

	if requestr.URL.Path != "/hello" {
		http.Error(responsew, "404 not found", http.StatusNotFound)

	}

	fmt.Fprint(responsew, "hello!")

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080 \n")

	err := http.ListenAndServe(":8080", nil) // this function always returns not nil error

	if err != nil {
		log.Fatal(err)
	}

}
