package router

import (
	"github.com/gorilla/mux"
	"github.com/gpunit2417/MongoAPI/controller"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/deletemovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}