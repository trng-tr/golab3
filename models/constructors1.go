package models

func NewCustomer(id string, firstname string, lastname string, email string, address Address,
	createdAt string, updatedAt string) (Customer, error) {
	if err := validateStrings(firstname, lastname, email, address.StreetName, address.ZipCode,
		address.City, address.Country); err != nil {
		return Customer{}, err
	}
	if err := validateNumber(float64(address.StreetNum)); err != nil {
		return Customer{}, err
	}
	return Customer{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Address:   address,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func NewAddress(id string, num int, street string, zip string, city string, country string) (Address, error) {
	if err := validateNumber(float64(num)); err != nil {
		return Address{}, err
	}
	if err := validateStrings(street, city, country, zip); err != nil {
		return Address{}, err
	}
	return Address{
		Id:         id,
		StreetNum:  num,
		StreetName: street,
		ZipCode:    zip,
		City:       city,
		Country:    country,
	}, nil
}

func NewOrderLine(productName string, quantity float64, unitPrice float64) (OrderLine, error) {
	if err := validateStrings(productName); err != nil {
		return OrderLine{}, err
	}

	if err := validateNumber(quantity, unitPrice); err != nil {
		return OrderLine{}, err
	}
	return OrderLine{
		ProductName: productName,
		Quantity:    quantity,
		UnitPrice:   unitPrice,
	}, nil
}

func NewOrder(id string, customer Customer, orderLines []OrderLine, createdAt string) (Order, error) {
	return Order{
		Id:         id,
		Customer:   customer,
		OrderLines: orderLines,
		CreatedAt:  createdAt,
	}, nil
}
