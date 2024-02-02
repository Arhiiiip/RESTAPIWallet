package main

import (
	handler "RESTAPIWallet/controllers"
	dbInterface "RESTAPIWallet/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := dbInterface.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	walletHandler := &handler.WalletHandler{DB: db}
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/wallet", walletHandler.CreateWallet).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/send", walletHandler.SendMoney).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/history", walletHandler.GetWalletHistory).Methods("GET")
	router.HandleFunc("/api/v1/wallet/{walletId}", walletHandler.GetWalletDetails).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
