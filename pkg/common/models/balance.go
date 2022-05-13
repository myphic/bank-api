package models

type Balance struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
