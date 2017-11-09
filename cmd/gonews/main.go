package main

import (
	"log"
	"net/http"
	"os"

	"github.com/naowal/gonews/pkg/app"
	"github.com/naowal/gonews/pkg/model"
)

const (
	port = ":8080"
	// mongoURL = "mongodb://127.0.0.1:27017"
	mongoURL = "mongodb://naowal:s00446204@ds251985.mlab.com:51985/gonews"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model: %v", err)
	}
	http.ListenAndServe(port, mux)
}
