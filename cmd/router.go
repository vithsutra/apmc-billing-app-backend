package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/handlers"
	"github.com/vsynclabs/billsoft/repository"
)

type Router struct {
	mux *mux.Router
	db *sql.DB
}

func NewRouter(conn *Connection) *Router {
	mux := mux.NewRouter()
	router := &Router{
		mux: mux,
		db: conn.db,
	}
	AuthRouters(router)
	return router
}

func AuthRouters(r *Router) {
	authHandler := handlers.NewUserHandler(repository.NewUserRepo(r.db))
	r.mux.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")
	r.mux.HandleFunc("/register", authHandler.RegisterHandler).Methods("POST")
}
