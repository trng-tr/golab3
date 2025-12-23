package service

import "github.com/trng-tr/golab3/model"

/* Le package des interfaces que les diffréntes structures
vont implementer
*/

type CustomerService interface {
	CreateCustomer(address model.Address) (model.Customer, error)
	PrintCustomer(c model.Customer)
}

type AddressService interface {
	CreateAddress() (model.Address, error)
}

type OrderLineService interface {
	CreateOrderLine(lineNumero int) (model.OrderLine, error)
}

type OrderService interface {
	CreateOrder(nbItems int, customer model.Customer) (model.Order, error)
	PrintOrder(o model.Order)
	AddOrderLineToOrder(line model.OrderLine) model.Order
}

type ProductService interface {
	CreateStockProduct(productNumero int) (model.Product, error)
	PrintStockProduct(p model.Product)
}
