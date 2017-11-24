package main

import (
	"io"
	"net/http"
)

//create a handler for Get
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

//create a mux that will keep a map of all the handlers.
var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	//variable server with Type http.Server
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{}, //an interface that is only required to implement one method whose signature is func(w http.ResponseWriter, r *http.Request)
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello //pass in handler we created and register it with a the route of the handler
	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}

//
// func main() {
// 	http.HandleFunc("/", hello)
// 	http.ListenAndServe(":8000", nil)
//
// }
