package models

import (
	"errors"
	"strings"

	"github.com/trng-tr/golab3/constantes"
)

/*
Un constructeur garantit que l’objet créé est valide, c'est pour cette raison
qu'en Go il est possible de mettre dans un constructeur des tests de validation
de l'objet pour éviter de créer un objet invalid
*/
func isStringEmpty(str string) bool {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}
func isStringShort(str string) bool {
	if len(strings.TrimSpace(str)) <= 0 {
		return true
	}
	return false
}

func validateStrings(strs ...string) error {
	for _, str := range strs {
		if isStringEmpty(str) {
			return errors.New(constantes.EmptyStrError)
		}
		if isStringShort(str) {
			return errors.New(constantes.ShortStrError)
		}
	}

	return nil
}

func isNumberValid(nbr float64) bool {
	return nbr > 0
}

func validateNumber(nbrs ...float64) error {
	for _, nbr := range nbrs {
		if !isNumberValid(nbr) {
			return errors.New(constantes.InvalidNbrError)
		}
	}

	return nil
}
