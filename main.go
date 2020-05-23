package main

import (
	"log"
	"net/http"

	"github.com/naspredam/rest-go-users/user"
)

func main() {
	var port = "8080"

	log.Println("Listening on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, user.Router()))
}
