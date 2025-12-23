package model

func NewProduct(id string, sku string, productName string, description string, price Price,
	stock Stock, createdAt string, updatedAt string, isActive bool) (Product, error) {

	if err := validateStrings(productName, description, price.Currency); err != nil {
		return Product{}, err
	}

	if err := validateNumber(price.UnitPrice, stock.Quantity); err != nil {
		return Product{}, err
	}

	return Product{
		Uuid:        id,
		Sku:         sku,
		ProductName: productName,
		Description: description,
		Price:       price,
		Stock:       stock,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		IsActive:    isActive,
	}, nil
}
