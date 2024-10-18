package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jakemackie/go-api/handlers"
)

func HandleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/videos", handlers.AllVideos).Methods("GET")
    myRouter.HandleFunc("/videos", handlers.PostVideo).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", myRouter))
}