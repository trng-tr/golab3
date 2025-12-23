package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"net/mail"

	"github.com/trng-tr/golab3/constantes"
	"github.com/trng-tr/golab3/reader"
	uuidpgk "github.com/trng-tr/golab3/uuid"
)

/*la structure Customer implemente la methode CreateCustomer()
de l'interface CustomerService*/

func (Customer) CreateCustomer(address Address) (Customer, error) {
	fmt.Print("Saisir le nom du client: ")
	lastname, err := reader.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return Customer{}, errors.New(constantes.ReadingError)
	}
	lastname = strings.TrimSpace(lastname)
	var uuid string = uuidpgk.GenerateUuid(lastname)

	fmt.Print("Saisir le prénom du client: ")
	firstname, err := reader.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return Customer{}, errors.New(constantes.ReadingError)
	}
	firstname = strings.TrimSpace(firstname)
	fmt.Print("Saisir le email  du client: ")
	email, err := reader.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return Customer{}, errors.New(constantes.ReadingError)
	}
	email = strings.TrimSpace(email)
	if !isValidEmail(email) {
		return Customer{}, errors.New("Email invalid")
	}
	var createdAt string = time.Now().Format("2006-01-01 15:01:06")
	c, err := NewCustomer(uuid, firstname, lastname, email, address, createdAt, "")
	if err != nil {
		return Customer{}, err
	}
	return c, nil
}

// packge Go net/mail pr valider un email
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
