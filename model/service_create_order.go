package model

import (
	"time"

	uuipkg "github.com/trng-tr/golab3/uuid"
)

/*
CreateOrder la structure Order implement la methode CreateOrder()
de l interface OrderService
*/
func (Order) CreateOrder(customer Customer, orderLines []OrderLine) (Order, error) {
	var uuid string = uuipkg.GenerateUuid("Order")
	// 2006-01-02 15:04:05 est la date de référence Go pour le formattage 👇
	var createdAt string = time.Now().Format("2006-01-02 15:04:05")

	order, err := NewOrder(uuid, customer, orderLines, createdAt)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}
