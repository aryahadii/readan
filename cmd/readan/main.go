package main

import (
	"log"
	"net/http"

	"github.com/aryahadii/readan/apis"
)

func main() {
	serve()
}

func serve() {
	router := apis.GetHTTPRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
