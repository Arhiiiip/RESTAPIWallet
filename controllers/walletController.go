package controllers

import (
	model "RESTAPIWallet/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type WalletHandler struct {
	DB *gorm.DB
}

func (walletHandler *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var walletID string

	for {
		walletID = uuid.New().String()
		var existingWallet model.Wallet
		result := walletHandler.DB.Where("id = ?", walletID).First(&existingWallet)
		if result.Error == gorm.ErrRecordNotFound {
			break
		} else if result.Error != nil {
			continue
		}
	}

	newWallet := model.Wallet{
		ID: walletID,
	}

	if err := walletHandler.DB.Create(&newWallet).Error; err != nil {
		http.Error(w, "Failed to create wallet", http.StatusInternalServerError)
		return
	}

	response := struct {
		ID      string  `json:"id"`
		Balance float64 `json:"balance"`
	}{
		ID:      newWallet.ID,
		Balance: newWallet.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (walletHandler *WalletHandler) SendMoney(w http.ResponseWriter, r *http.Request) {
	//todo Реализовать логику перевода средств
}

func (walletHandler *WalletHandler) GetWalletHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	walletID := params["walletId"]
	var transactions []model.Transaction
	if err := walletHandler.DB.Where("from = ? OR to = ?", walletID, walletID).Find(&transactions).Error; err != nil {
		http.Error(w, "Failed to retrieve transaction history", http.StatusInternalServerError)
		return
	}

	response := make([]struct {
		Time   time.Time `json:"time"`
		From   string    `json:"from"`
		To     string    `json:"to"`
		Amount float64   `json:"amount"`
	}, len(transactions))

	for i, transaction := range transactions {
		response[i] = struct {
			Time   time.Time `json:"time"`
			From   string    `json:"from"`
			To     string    `json:"to"`
			Amount float64   `json:"amount"`
		}{
			Time:   transaction.Time,
			From:   transaction.From,
			To:     transaction.To,
			Amount: transaction.Amount,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (walletHandler *WalletHandler) GetWalletDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	walletID := params["walletId"]
	var wallet model.Wallet
	if err := walletHandler.DB.First(&wallet, "id = ?", walletID).Error; err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	response := struct {
		ID      string  `json:"id"`
		Balance float64 `json:"balance"`
	}{
		ID:      wallet.ID,
		Balance: wallet.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
