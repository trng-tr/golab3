package model

type Customer struct {
	Id        string
	Firstname string
	Lastname  string
	Email     string
	Address   Address
	CreatedAt string //date and time
	UpdatedAt string //date and time
}

type Address struct {
	Id         string
	StreetNum  int
	StreetName string
	ZipCode    string
	City       string
	Country    string
}

type OrderLine struct {
	ProductName string
	Quantity    float64
	UnitPrice   float64
}

type Order struct {
	Id         string
	Customer   Customer
	OrderLines []OrderLine
	CreatedAt  string
}

type Product struct { // product to store
	Uuid        string
	Sku         string // référence du produit
	ProductName string
	Description string
	Price       Price
	Stock       Stock
	CreatedAt   string // date and time
	UpdatedAt   string // date and time
	IsActive    bool
}

type Price struct {
	UnitPrice float64 // le prix peut changer avec le temps
	Currency  string
}

type Stock struct {
	Quantity float64 // qté du produit en stock
}

const (
	Euro   string = "€"
	Dollar string = "$"
)
