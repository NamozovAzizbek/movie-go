package routes

import (
	"net/http"
	"./controllers"
)

var RegisterMovie = func ()  {
	http.HandleFunc("/movies", controllers.GetBook)
}
