package models

type Products struct {
	ID          int     `json:"id" gorm:"primary_key"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}
