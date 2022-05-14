package models

type Balance struct {
	Id       int     `json:"id" gorm:"primaryKey"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
