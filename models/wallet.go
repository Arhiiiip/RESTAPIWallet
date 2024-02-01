package models

type Wallet struct {
	ID      string  `gorm:"primaryKey"`
	Balance float64 `gorm:"default:100"`
}
