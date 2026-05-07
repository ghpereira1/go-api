package model

type Product struct {
	ID int `json:"id"`
	// binding:"required" impede nomes vazios
	Name string `json:"product_name" binding:"required"`
	// binding:"required,gt=0" obriga o campo e exige que seja Greater Than (maior que) 0
	Price float64 `json:"price" binding:"required,gt=0"`
}
