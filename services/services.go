package services

import "github.com/trng-tr/golab3/models"

/* Le package des interfaces que les diffréntes structures
vont implementer
*/

type CustomerService interface {
	CreateCustomer(address models.Address) (models.Customer, error)
	PrintCustomer(c models.Customer)
}

type AddressService interface {
	CreateAddress() (models.Address, error)
}

type OrderLineService interface {
	CreateOrderLine(lineNumero int) (models.OrderLine, error)
}

type OrderService interface {
	CreateOrder(nbItems int, customer models.Customer) (models.Order, error)
	PrintOrder(o models.Order)
}

type ProductService interface {
	CreateStockProduct(productNumero int) (models.Product, error)
	PrintStockProduct(p models.Product)
}
