package models

import (
	"time"
)

type Transaction struct {
	Time    time.Time
	From    string
	To      string
	Amount  float64
	Wallet  Wallet `gorm:"foreignKey:From"`
	Wallet2 Wallet `gorm:"foreignKey:To"`
}
