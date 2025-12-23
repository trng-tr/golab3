package models

/*
 methode AddOrderLineInOrder de la structure
*/
func (o *Order) AddOrderLineToOrder(orderLine OrderLine) {
	o.OrderLines = append(o.OrderLines, orderLine)
}
