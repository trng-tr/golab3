package main

import (
	"fmt"

	"github.com/trng-tr/golab3/model"
	"github.com/trng-tr/golab3/service"
)

func main() {
	//la structure address implément interface AddresseService👇
	var addressSvc service.AddressService = model.Address{}
	address, err := addressSvc.CreateAddress()
	//la structure Customer implement interface CustomerService👇
	var c service.CustomerService = model.Customer{}
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

	stock, err := service.FillStock(nb)
	if err != nil {
		fmt.Println(err)
		return
	}
	//la structure Product implement interface ProductService👇
	var pdtService service.ProductService = model.Product{}
	for _, product := range stock {
		pdtService.PrintStockProduct(product)
	}
	//la structure Order implement interface OrderService👇
	var orderSvc = model.Order{}
	order, err := orderSvc.CreateOrder(nb, customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	orderSvc.PrintOrder(order)
	//la structure OrderLine implement interface OrderLineService👇
	var orderLineSvc service.OrderLineService = model.OrderLine{}
	orderLine, err := orderLineSvc.CreateOrderLine(len(order.OrderLines) + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	// la structure Order define la méthode AddOrderLineToOrder👇
	// mais ce n'est pas une méthode d interface
	var orderNew = order.AddOrderLineToOrder(orderLine)
	orderSvc.PrintOrder(orderNew)

}
