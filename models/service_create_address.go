package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/trng-tr/golab3/constantes"
	"github.com/trng-tr/golab3/reader"
	"github.com/trng-tr/golab3/uuid"
)

/*
la structure Address implemente la methode CreateAddress
de l'interface AddressService
*/
func (Address) CreateAddress() (Address, error) {
	fmt.Print("Saisir le numero de la maison: ")
	str, err := reader.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return Address{}, errors.New(constantes.ReadingError)
	}

	str = strings.TrimSpace(str)
	streetNum, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return Address{}, errors.New(constantes.ConversionError)
	}

	fmt.Print("Saisir le nom de la rue: ")
	streetName, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Address{}, errors.New(constantes.ReadingError)
	}
	streetName = strings.TrimSpace(streetName)

	var id string = uuid.GenerateUuid((streetName))

	fmt.Print("Saisir la boite postale: ")
	zipCode, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Address{}, errors.New(constantes.ReadingError)
	}

	zipCode = strings.TrimSpace(zipCode)

	fmt.Print("Saisir le nom de la ville: ")
	city, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Address{}, errors.New(constantes.ReadingError)
	}
	city = strings.TrimSpace(city)

	fmt.Print("Saisir le nom du pays: ")
	country, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return Address{}, errors.New(constantes.ReadingError)
	}
	country = strings.TrimSpace(country)
	address, err := NewAddress(id, int(streetNum), streetName, zipCode, city, country)
	if err != nil {
		return Address{}, err
	}

	return address, nil
}
