package model

import "fmt"

// PrintStockProduct /*la structure Product implement la methode PrintStockProduct()
// de interface ProductService*/
func (Product) PrintStockProduct(p Product) {
	fmt.Println("{")
	fmt.Printf("    Id: %s\n", p.Uuid)
	fmt.Printf("    Sku: %s\n", p.Sku)
	fmt.Printf("    Prod Name: %s\n", p.ProductName)
	fmt.Printf("    Description: %s\n", p.Description)
	fmt.Println("    Price: {")
	fmt.Printf("      PU: %.2f\n", p.Price.UnitPrice)
	fmt.Printf("      Currency: %s\n", p.Price.Currency)
	fmt.Println("   }")
	fmt.Println("   Stock: {")
	fmt.Printf("      Qté stock: %.2f\n", p.Stock.Quantity)
	fmt.Println("   }")
	fmt.Printf("   PT produit: %.2f %s\n", p.Price.UnitPrice*p.Stock.Quantity, p.Price.Currency)
	fmt.Printf("   Date création: %s\n", p.CreatedAt)
	fmt.Printf("   Date MJ: %s\n", p.UpdatedAt)
	fmt.Printf("   Actif: %v\n", p.IsActive)
	fmt.Println("}")
}
