package routers

import (
	"handlers"

	"github.com/gorilla/mux"
)

func setBoardRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/boards/auth/{ownerid}", handlers.GetAllBoardsEndpoint).Methods("GET")
	router.HandleFunc("/api/boards/{id}/auth/{ownerid}", handlers.GetBoardEndpoint).Methods("GET")
	router.HandleFunc("/api/boards/auth/{ownerid}", handlers.CreateBoardEndpoint).Methods("POST")
	router.HandleFunc("/api/boards/auth/{ownerid}", handlers.DeleteBoardEndpoint).Methods("DELETE")
	router.HandleFunc("/api/boards/{id}/auth/{ownerid}", handlers.DeleteBoardByIDEndpoint).Methods("DELETE")
	router.HandleFunc("/api/boards/{id}/auth/{ownerid}", handlers.UpdateBoardEndpoint).Methods("PUT")
	//router.HandleFunc("/boards/{id}/categories/{id2}/auth/{ownerid}", handlers.CreateJobEndpoint).Methods("PUT") // This should not have to be used

	return router
}
