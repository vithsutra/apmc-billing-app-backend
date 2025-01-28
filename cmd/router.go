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
	ConsigneeRouters(router)
	ReceiverRouters(router)
	ProductRouters(router)
	return router
}

func UserRouters(r *Router) {
	userHandler := handlers.NewUserHandler(repository.NewUserRepo(r.db))
	r.mux.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
	r.mux.HandleFunc("/register", userHandler.RegisterHandler).Methods("POST")
	r.mux.HandleFunc("/delete/{user_id}", userHandler.DeleteUserHandler).Methods("DELETE")
}

func ConsigneeRouters(r *Router) {
	ConsigneeHandler := handlers.NewConsigneeHandler(repository.NeeConsigneeRepo(r.db))
	r.mux.HandleFunc("/create/consignee", ConsigneeHandler.CreateConsigneeHandler).Methods("POST")
	r.mux.HandleFunc("/delete/consignee/{consignee_id}", ConsigneeHandler.DeleteConsigneeHandler).Methods("DELETE")
	r.mux.HandleFunc("/get/consignees/{user_id}", ConsigneeHandler.GetConsigneeHandler).Methods("GET")
}

func ReceiverRouters(r *Router) {
	receiverHandler := handlers.NewReceiverHandler(repository.NewReceiverRepo(r.db))
	r.mux.HandleFunc("/create/receiver", receiverHandler.CreateReceiverHandler).Methods("POST")
	r.mux.HandleFunc("/delete/receiver/{receiver_id}", receiverHandler.DeleteReceiverHandler).Methods("DELETE")
	r.mux.HandleFunc("get/receivers/{user_id}", receiverHandler.GetReceiversHandler).Methods("GET")
}

func ProductRouters(r *Router) {
	productHandler := handlers.NewProductHandler(repository.NewProductRepo(r.db))
	r.mux.HandleFunc("/create/product", productHandler.CreateProductHandler).Methods("POST")
	r.mux.HandleFunc("/delete/product/{product_id}", productHandler.DeleteProductHandler).Methods("DELETE")
}
