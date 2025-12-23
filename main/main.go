package main

import (
	"fmt"

	"github.com/trng-tr/golab3/models"
	"github.com/trng-tr/golab3/services"
)

func main() {
	//la structure addresse implemente l'interface AddresseService👇
	var addressSvc services.AddressService = models.Address{}
	address, err := addressSvc.CreateAddress()
	//la structure Customer implemente l'interface CustomerService👇
	var c services.CustomerService = models.Customer{}
	customer, err := c.CreateCustomer(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.PrintCustomer(customer)

	nb, err := inputNbr()
	if err != nil {
		fmt.Println(err)
		return
	}

	stock, err := services.FillStock(nb)
	if err != nil {
		fmt.Println(err)
		return
	}
	//la structure Product implemente l'interface ProductService👇
	var pdtService services.ProductService = models.Product{}
	for _, product := range stock {
		pdtService.PrintStockProduct(product)
	}
	//la structure Order implemente l'interface OrderService👇
	var orderSvc = &models.Order{}
	order, err := orderSvc.CreateOrder(nb, customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	orderSvc.PrintOrder(order)
	//la structure OrderLine implemente l'interface OrderLineService👇
	var orderLineSvc services.OrderLineService = models.OrderLine{}
	orderLine, err := orderLineSvc.CreateOrderLine(len(order.OrderLines) + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	// la structure Order définit la méthode AddOrderLineToOrder👇
	// mais ce n'est pas une méthode d'interface
	order.AddOrderLineToOrder(orderLine)
	orderSvc.PrintOrder(order)

}
