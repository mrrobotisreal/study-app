package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"study-app/db_stuff"
)

var Mux *http.ServeMux

func GetFart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Faaarrrrrttttttttt!!!!!!!!!")
	io.WriteString(w, "Faaarrrrrttttttttt!!!!!!!!!")
}

func GetCollections(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-type")
	response.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	response.Header().Add("content-type", "application/json")
	type GetCollectionsRequest struct {
		user string
	}
	var user GetCollectionsRequest
	json.NewDecoder(request.Body).Decode(&user)
	collections := db_stuff.GetCollections(user.user)
	json.NewEncoder(response).Encode(&collections)
}

func GetCollection(response http.ResponseWriter, request *http.Request) {
	urlParams := strings.Split(request.URL.Path, "/")
	for i := 0; i < len(urlParams); i++ {
		fmt.Println("The current url param be like:")
		fmt.Println(urlParams[i])
	}
}
