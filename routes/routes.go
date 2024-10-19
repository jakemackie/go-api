package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jakemackie/go-api/controllers"
)

func HandleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/videos", controllers.AllVideos).Methods("GET")
    myRouter.HandleFunc("/videos", controllers.PostVideo).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", myRouter))
}