package router

import (
	"github.com/gorilla/mux"
	controller "github.com/shahanahmed86/mongo-api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controller.GetMovie).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteMovies).Methods("DELETE")

	return router
}
