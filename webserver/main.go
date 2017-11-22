package main

import (
	"fmt"
	"net/http"
	"strings"
)

//create handler

func sayHello(w http.ResponseWriter, r *http.Request) {
	//retrieve the path of the URL,
	message := r.URL.Path
	//removes the first slash
	message = strings.TrimPrefix(message, "/")
	//append hello to the beginning of the sentence
	message = "Hello " + message
	//write the final message to the ResponseWriter converted to bytes
	w.Write([]byte(message))
	fmt.Println("We have a response :) ")
}

func main() {
	//use handler sayHello to every request that hits the server
	http.HandleFunc("/", sayHello)
	//start the server on port 8080 and handle an error if something wrong happens
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// // create a webserver
//
// //this creates a handler, which retrieves the path of the URL
// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello World") //a helper function to let you write a string into a given writable stream
// }
//
// func main() {
// 	http.HandleFunc("/", hello) //register another function to be the handle function (the hello function) - route you want to match,
// 	http.ListenAndServe(":8000", nil)
//
// }
