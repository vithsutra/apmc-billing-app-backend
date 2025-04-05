package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/handlers"
	"github.com/vsynclabs/billsoft/pkg/storage"
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
	InvoiceRouters(router)
	BillerRouters(router.mux, router.db)
	BankerRouter(router)
	return router
}
func UserRouters(r *Router) {
	userHandler := handlers.NewUserHandler(repository.NewUserRepo(r.db))
	r.mux.HandleFunc("/create/user", userHandler.RegisterHandler).Methods("POST")
	r.mux.HandleFunc("/delete/{user_id}", userHandler.DeleteUserHandler).Methods("DELETE")
	r.mux.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")
	r.mux.HandleFunc("/forgotpassword", userHandler.ForgotPasswordHandler).Methods("POST")
	r.mux.HandleFunc("/validateotp", userHandler.ValidateOTPHandler).Methods("POST")
	r.mux.HandleFunc("/resetpassword", userHandler.ResetPasswordHandler).Methods("POST")
}

func ConsigneeRouters(r *Router) {
	ConsigneeHandler := handlers.NewConsigneeHandler(repository.NewConsigneeRepo(r.db))
	r.mux.HandleFunc("/create/consignee", ConsigneeHandler.CreateConsigneeHandler).Methods("POST")
	r.mux.HandleFunc("/delete/consignee/{consignee_id}", ConsigneeHandler.DeleteConsigneeHandler).Methods("DELETE")
	r.mux.HandleFunc("/get/consignees/{user_id}", ConsigneeHandler.GetConsigneeHandler).Methods("GET")
}

func ReceiverRouters(r *Router) {
	receiverHandler := handlers.NewReceiverHandler(repository.NewReceiverRepo(r.db))
	r.mux.HandleFunc("/create/receiver", receiverHandler.CreateReceiverHandler).Methods("POST")
	r.mux.HandleFunc("/delete/receiver/{receiver_id}", receiverHandler.DeleteReceiverHandler).Methods("DELETE")
	r.mux.HandleFunc("/get/receivers/{user_id}", receiverHandler.GetReceiversHandler).Methods("GET")
}

func ProductRouters(r *Router) {
	productHandler := handlers.NewProductHandler(repository.NewProductRepo(r.db))
	r.mux.HandleFunc("/create/product", productHandler.CreateProductHandler).Methods("POST")
	r.mux.HandleFunc("/delete/product/{product_id}", productHandler.DeleteProductHandler).Methods("DELETE")
	r.mux.HandleFunc("/get/products/{invoice_id}", productHandler.GetProductHandler).Methods("GET")
}

func InvoiceRouters(r *Router) {
	invoiceHandler := handlers.NewInvoiceHandler(repository.NewInvoiceRepo(r.db))
	r.mux.HandleFunc("/create/invoice", invoiceHandler.CreateInvoiceHandler).Methods("POST")
	r.mux.HandleFunc("/delete/invoice/{invoice_id}", invoiceHandler.DeleteInvoiceHandler).Methods("DELETE")
	r.mux.HandleFunc("/get/invoices/{user_id}", invoiceHandler.GetInvoicesHandler).Methods("GET")
	r.mux.HandleFunc("/update/invoice/payment/{invoice_id}", invoiceHandler.UpdateInvoicePaymentStatusHandler).Methods("PATCH")
	r.mux.HandleFunc("/download/invoice/{invoice_id}", invoiceHandler.DownloadInvoiceHandler).Methods("GET")
}

func BillerRouters(r *mux.Router, db *sql.DB) {
	s3Repo := &storage.AwsS3Repo{}
	billerHandler := handlers.NewBillerHandler(repository.NewBillerRepo(db, s3Repo))

	r.HandleFunc("/create/biller", billerHandler.CreateBillerHandler).Methods("POST")
	r.HandleFunc("/delete/biller/{biller_id}", billerHandler.DeleteBillerHandler).Methods("DELETE")
	r.HandleFunc("/get/billers/{user_id}", billerHandler.GetBillerHandler).Methods("GET")
	r.HandleFunc("/upload/company/logo/{userId}", billerHandler.UploadCompanyLogoHandler).Methods("POST")
	r.HandleFunc("/delete/company/logo/{fileName}", billerHandler.DeleteCompanyLogoHandler).Methods("DELETE")
}

func BankerRouter(r *Router) {
	bankerHandler := handlers.NewBankerHandler(repository.NewBankerRepo(r.db))
	r.mux.HandleFunc("/create/banker", bankerHandler.CreateBankerHandler).Methods("POST")
	r.mux.HandleFunc("/delete/banker/{banker_id}", bankerHandler.DeleteBankerHandler).Methods("DELETE")
	r.mux.HandleFunc("/get/bankers/{user_id}", bankerHandler.GetBankerHandler).Methods("GET")
}
