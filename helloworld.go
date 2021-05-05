package main

import (
	"net/http"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")
	res.Write(data)
}

func main() {
	// create a new handler
	handler := HttpHandler{} // listen and serve
	http.ListenAndServe(":9000", handler)
}
