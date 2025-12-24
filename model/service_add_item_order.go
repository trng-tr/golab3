package model

// AddOrderLineToOrder
// la structure Order implement l interface OrderService
func (o Order) AddOrderLineToOrder(orderLine OrderLine) Order {
	o.OrderLines = append(o.OrderLines, orderLine)
	return o
}
