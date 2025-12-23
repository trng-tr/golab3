package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/trng-tr/golab3/constantes"
	"github.com/trng-tr/golab3/reader"
)

/*
la structure OrderLine implemente la methode CreateOrderLine()
de l'interface OrderLineSrvice
*/
func (OrderLine) CreateOrderLine(numero int) (OrderLine, error) {
	fmt.Printf("saisir le nom du produit  %d à commander: ", numero)
	productName, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return OrderLine{}, errors.New(constantes.ReadingError)
	}
	productName = strings.TrimSpace(productName)
	fmt.Printf("saisir la quantité pour le produit  %d a commander: ", numero)
	strQte, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return OrderLine{}, errors.New(constantes.ReadingError)
	}
	strQte = strings.TrimSpace(strQte)
	if strQte == "" {
		return OrderLine{}, errors.New(constantes.EmptyStrError)
	}
	qte, err := strconv.ParseFloat(strQte, 64)
	if err != nil {
		return OrderLine{}, errors.New(constantes.ConversionError)
	}
	fmt.Printf("saisir le prix pour le produit  %d a commander: ", numero)
	strPrice, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return OrderLine{}, errors.New(constantes.ReadingError)
	}
	strPrice = strings.TrimSpace(strPrice)

	if strPrice == "" {
		return OrderLine{}, errors.New(constantes.EmptyStrError)
	}
	price, err := strconv.ParseFloat(strPrice, 64)
	if err != nil {
		return OrderLine{}, errors.New(constantes.ConversionError)
	}
	orderLine, err := NewOrderLine(productName, qte, price)
	if err != nil {
		return OrderLine{}, err
	}
	return orderLine, nil
}
