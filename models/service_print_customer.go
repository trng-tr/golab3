package models

import "fmt"

/*la structure Customer implemente la methode PrintCustomer()
de l'interface CustomerService*/
func (Customer) PrintCustomer(c Customer) {
	fmt.Println("{")
	fmt.Println(" customer {")
	fmt.Printf("     id: %s\n", c.Id)
	fmt.Printf("     prénom: %s\n", c.Firstname)
	fmt.Printf("     nom: %s\n", c.Lastname)
	fmt.Printf("     email: %s\n", c.Email)
	fmt.Println("     address {")
	fmt.Printf("       id: %s\n", c.Address.Id)
	fmt.Printf("       numero: %d\n", c.Address.StreetNum)
	fmt.Printf("       nom de la rue: %s\n", c.Address.StreetName)
	fmt.Printf("       code postal: %s\n", c.Address.ZipCode)
	fmt.Printf("       ville: %s\n", c.Address.City)
	fmt.Printf("       pays: %s\n", c.Address.Country)
	fmt.Println("    }")
	fmt.Println("  }")
	fmt.Println("}")
}
