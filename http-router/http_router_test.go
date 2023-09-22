package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHTTPRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Router")
	})

	request := httptest.NewRequest("GET", "http://localhost:5137/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	reponse := recorder.Result()

	body, _ := io.ReadAll(reponse.Body)

	// fmt.Print(string(body))
	assert.Equal(t, "Hello Router", string(body))
}

func TestHTTPRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprintln(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:5137/products/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	reponse := recorder.Result()

	body, _ := io.ReadAll(reponse.Body)

	fmt.Print(string(body))
	// assert.Equal(t, "Hello Router", string(body))
}

func TestHTTPRouterPatternNamedParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " items " + itemId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:5137/products/2/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	reponse := recorder.Result()

	body, _ := io.ReadAll(reponse.Body)

	fmt.Print(string(body))
	assert.Equal(t, "Product 2 items 1", string(body))
}
