package service

import (
	"errors"
	"fmt"

	"github.com/trng-tr/golab3/model"
)

// methde indépendante des structures
func FillStock(nbProducts int) ([]model.Product, error) {
	if nbProducts <= 0 {
		return nil, errors.New("nombre de produits pour le stock est invalid")
	}
	var stockProducts []model.Product = make([]model.Product, 0, nbProducts)
	/*la structure Produc implemente la methode CreateStockProduct()
	de l'interface ProductService*/
	var prdService = model.Product{}
	for i := 1; i <= nbProducts; i++ {
		product, err := prdService.CreateStockProduct(i)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		stockProducts = append(stockProducts, product)
	}
	return stockProducts, nil
}
