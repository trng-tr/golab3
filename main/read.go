package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/trng-tr/golab3/constantes"
	"github.com/trng-tr/golab3/reader"
)

func inputNbr() (int, error) {
	fmt.Print("saisir le nombre de produits: ")
	str, err := reader.StreamReader.ReadString('\n')
	if err != nil {
		return -1, errors.New(constantes.ReadingError)
	}
	str = strings.TrimSpace(str)
	if str == "" {
		return -1, errors.New(constantes.EmptyStrError)
	}

	nbItems, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return -1, errors.New(constantes.ConversionError)
	}

	return int(nbItems), nil
}
