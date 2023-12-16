package model

type Product struct {
	ID       int     `json:"ID"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Category string  `json:"category"`
}
