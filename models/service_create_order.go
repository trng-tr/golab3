package models

import (
	"errors"
	"time"

	uuipkg "github.com/trng-tr/golab3/uuid"
)

/*
la structure Order implemente la methode CreateOrder()
de l'interface OrderService
*/
func (Order) CreateOrder(nbItems int, customer Customer) (Order, error) {
	if nbItems <= 0 {
		return Order{}, errors.New("nompre de produits pour cette cmd invalid")
	}
	var uuid string = uuipkg.GenerateUuid("Order")
	var ordeLines []OrderLine = make([]OrderLine, 0, nbItems)
	// la struture OrderLine implemente l'interface OrderLineService👇
	var orderLineSvc = OrderLine{}

	for i := 1; i <= nbItems; i++ {
		orderLine, err := orderLineSvc.CreateOrderLine(i)
		if err != nil {
			return Order{}, err
		}
		ordeLines = append(ordeLines, orderLine)
	}
	// date de référence Go pour le formattage 👇
	var createdAt string = time.Now().Format("2006-01-02 15:04:05")

	return NewOrder(uuid, customer, ordeLines, createdAt)
}
