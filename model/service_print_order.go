package model

import "fmt"

//PrintOrder
/*
la structure Order implemente la methode PrintOrder()
de l'interface OrderService
*/
func (Order) PrintOrder(o Order) {
	fmt.Println("{")
	fmt.Printf("  id: %s\n", o.Id)
	fmt.Println("  customer {")
	fmt.Printf("     id: %s\n", o.Customer.Id)
	fmt.Printf("     prénom: %s\n", o.Customer.Firstname)
	fmt.Printf("     nom: %s\n", o.Customer.Lastname)
	fmt.Printf("     email: %s\n", o.Customer.Email)
	fmt.Println("    address {")
	fmt.Printf("       id: %s\n", o.Customer.Address.Id)
	fmt.Printf("       numero: %d\n", o.Customer.Address.StreetNum)
	fmt.Printf("       nom de la rue: %s\n", o.Customer.Address.StreetName)
	fmt.Printf("       code postal: %s\n", o.Customer.Address.ZipCode)
	fmt.Printf("       ville: %s\n", o.Customer.Address.City)
	fmt.Printf("       pays: %s\n", o.Customer.Address.Country)
	fmt.Println("    }")
	fmt.Println("  }")
	fmt.Println("  items: [")
	for i, item := range o.OrderLines {
		fmt.Printf("   %d : libelé: %s, quantité: %.2f, pu: %.2f, pt: %.2f\n",
			i+1,
			item.ProductName,
			item.Quantity,
			item.UnitPrice,
			item.UnitPrice*float64(item.Quantity),
		)
	}
	fmt.Println("  ]")

	fmt.Printf("  orderDate: %s\n", o.CreatedAt)
	fmt.Println("}")
}
