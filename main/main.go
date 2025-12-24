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

	var ordeLines []model.OrderLine = make([]model.OrderLine, 0, nb)
	// la struture OrderLine implemente l'interface OrderLineService👇
	var orderLineSvc service.OrderLineService = model.OrderLine{}
	for i := 1; i <= nb; i++ {
		orderLine, err := orderLineSvc.CreateOrderLine(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		ordeLines = append(ordeLines, orderLine)
	}

	//la structure Order implement interface OrderService👇
	var orderSvc = model.Order{}
	order, err := orderSvc.CreateOrder(nb, customer, ordeLines)
	if err != nil {
		fmt.Println(err)
		return
	}
	orderSvc.PrintOrder(order)
	// une nouvelle orderLine
	orderLine, err := orderLineSvc.CreateOrderLine(len(order.OrderLines) + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	var orderNew = order.AddOrderLineToOrder(orderLine)
	orderSvc.PrintOrder(orderNew)
}
