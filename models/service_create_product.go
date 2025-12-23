package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/trng-tr/golab3/constantes"
	"github.com/trng-tr/golab3/reader"
	uuipkg "github.com/trng-tr/golab3/uuid"
)

/*
la structure Product implemente la methode CreateStockProduct()
de l'interface ProductService
*/
func (Product) CreateStockProduct(productNumero int) (Product, error) {
	fmt.Printf("saisir le nom du produit %d: ", productNumero)
	prodName, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Product{}, errors.New(constantes.ReadingError)
	}
	prodName = strings.TrimSpace(prodName)
	var uuid string = uuipkg.GenerateUuid(prodName)
	var sku string = uuipkg.GenerateUuid(prodName)

	fmt.Printf("saisir la description du produit %d:", productNumero)
	desc, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Product{}, errors.New(constantes.ReadingError)
	}
	desc = strings.TrimSpace(desc)
	fmt.Printf("saisir le prix pour le produit %d: ", productNumero)
	strPrice, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Product{}, errors.New(constantes.ReadingError)
	}
	strPrice = strings.TrimSpace(strPrice)
	price, err := strconv.ParseFloat(strPrice, 64)
	if err != nil {
		return Product{}, errors.New(constantes.ConversionError)
	}

	fmt.Printf("saisir le currency pour le prix du produit %d: ", productNumero)
	currency, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Product{}, errors.New(constantes.ReadingError)
	}
	currency = strings.TrimSpace(currency)
	if currency != Euro && currency != Dollar {
		return Product{}, errors.New("error: unknown currency")
	}
	var objPrice Price = Price{
		UnitPrice: price,
		Currency:  currency,
	}
	fmt.Printf("saisir la quantité pour le produit%d: ", productNumero)
	str, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Product{}, errors.New(constantes.ReadingError)
	}
	str = strings.TrimSpace(str)
	qte, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return Product{}, errors.New(constantes.ConversionError)
	}
	var stock Stock = Stock{
		Quantity: qte,
	}

	product, err := NewProduct(uuid, sku, prodName, desc, objPrice, stock,
		time.Now().Format("2006-01-02 15:04:05"), "", true) /*2006-01-02 15:04:05 est la
	date de référence en Go*/
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
