package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello HTTP Router")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:5137",
	}

	server.ListenAndServe()
}
