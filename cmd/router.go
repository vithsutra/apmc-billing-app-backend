package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/handlers"
	"github.com/vsynclabs/billsoft/repository"
)

type Router struct {
	mux *mux.Router
	db  *sql.DB
}

func NewRouter(conn *Connection) *Router {
	mux := mux.NewRouter()
	router := &Router{
		mux: mux,
		db:  conn.db,
	}
	UserRouters(router)
	return router
}

func UserRouters(r *Router) {
	userHandler := handlers.NewUserHandler(repository.NewUserRepo(r.db))
	r.mux.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
	r.mux.HandleFunc("/register", userHandler.RegisterHandler).Methods("POST")
	r.mux.HandleFunc("/delete/{user_id}", userHandler.DeleteUserHandler).Methods("DELETE")
}
