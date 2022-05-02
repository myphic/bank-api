package models

type Balance struct {
	Id     int   `json:"id" gorm:"primaryKey"`
	Amount int32 `json:"amount"`
}
