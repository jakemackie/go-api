package main

import (
	"github.com/jakemackie/go-api/handlers"
	"github.com/jakemackie/go-api/models"
	"github.com/jakemackie/go-api/routes"
)

func main() {
    // Initialise the videos array to keep track of all videos in memory
    handlers.Videos = models.Videos{
        {
            Title:       "Example 1 Minute Video",
            Description: "An example of a 1 minute video",
            Length:      60,
        },
    }
    routes.HandleRequests()
}