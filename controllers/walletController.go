package controllers

import (
	model "RESTAPIWallet/models"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
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
	//todo Разработать логику возврата истории транзакций
}

func (walletHandler *WalletHandler) GetWalletDetails(w http.ResponseWriter, r *http.Request) {
	//todo Разрботать логику возврата детайлей кошелька
}
