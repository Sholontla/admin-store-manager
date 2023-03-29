package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProviderInformation struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	ProviderBuissness string             `bson:"provider_buissness" json:"provider_buissness"`
	ProviderContacts  []ProviderContact  `bson:"provider_contact" json:"provider_contact"`
	ProviderAddress   []ProviderAddress  `bson:"provider_address" json:"provider_address"`
	ProviderProducts  []ProviderProducts `bson:"provider_products" json:"provider_products"`
	CreatedAt         string             `bson:"created_at" json:"created_at"`
	UpdatedAt         string             `bson:"updated_at" json:"updated_at"`
	AdminUser         string             `bson:"admin_user" json:"admin_user"`
}

type ProviderContact struct {
	ContactName          string `bson:"contact_name" json:"contact_name"`
	ContactLastName      string `bson:"contact_last_name" json:"contact_last_name"`
	ContactPhoneNumber   string `bson:"contact_phone_number" json:"contact_phone_number"`
	ContactMessageNumber string `bson:"contact_message_number" json:"contact_message_number"`
	ContactEmail         string `bson:"contact_email" json:"contact_email"`
}

type ProviderAddress struct {
	ProviderStreet  string `bson:"provider_street" json:"provider_street"`
	ProviderNumber  string `bson:"provider_number" json:"provider_number"`
	ProviderZipCode string `bson:"provider_zip_code" json:"provider_zip_code"`
	ProviderCity    string `bson:"provider_city" json:"provider_city"`
}

type ProviderProducts struct {
	ProductName           string  `bson:"product_name" json:"product_name"`
	ProductCategory       string  `bson:"product_ctegory" json:"product_ctegory"`
	ProductPrice          float64 `bson:"product_price" json:"product_price"`
	ProductSKU            string  `bson:"product_sku" json:"product_sku"`
	ProductMaterial       string  `bson:"product_material" json:"product_material"`
	ProductClassification string  `bson:"product_classification" json:"product_classification"`
	ProductQuantity       int     `bson:"product_quantity" json:"product_quantity"`
}

func (p ProviderInformation) Validations(ProviderInformation) ProviderInformation {

	for i := range p.ProviderProducts {
		switch {
		case p.ProviderProducts[i].ProductPrice >= 0.01 && p.ProviderProducts[i].ProductPrice <= 15.00:
			p.ProviderProducts[i].ProductClassification = "A"

		case p.ProviderProducts[i].ProductPrice >= 15.01 && p.ProviderProducts[i].ProductPrice <= 20.00:
			p.ProviderProducts[i].ProductClassification = "B"

		case p.ProviderProducts[i].ProductPrice >= 20.01 && p.ProviderProducts[i].ProductPrice <= 30.00:
			p.ProviderProducts[i].ProductClassification = "C"

		case p.ProviderProducts[i].ProductPrice >= 30.01:
			p.ProviderProducts[i].ProductClassification = "D"

		}
		for c := range p.ProviderContacts {
			var (
				pro = strings.ToUpper(p.ProviderContacts[c].ContactName)
				pn  = SelectStringElement(strings.ToUpper(p.ProviderProducts[i].ProductName), 5)
				pca = SelectStringElement(strings.ToUpper(p.ProviderProducts[i].ProductCategory), 5)
				pc  = SelectStringElement(strings.ToUpper(p.ProviderProducts[i].ProductClassification), 1)
				pm  = SelectStringElement(p.ProviderProducts[i].ProductMaterial, 10)

				r = fmt.Sprintf(pro + "-" + pn + "-" + pca + "-" + pc + "-" + pm + "-" + uuid.New().String())
			)

			p.ProviderProducts[i].ProductSKU = r
		}

	}

	return p
}

// func (a *ProviderInformation) CalculateTotalPrice() *float64 {

// 	var revenue float64 = 0.0

// 	for _, orderItem := range a.SupplierProducts {
// 		revenue += orderItem.ProductPrice

// 	}

// 	return &revenue

// }

// func isEmailValid(e string) bool {
// 	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	return emailRegex.MatchString(e)
// }

// func (customer *Supplier) UserValidation() utilerrors.RestErr {
// 	customer.SupplierEmailNumber = strings.TrimSpace(strings.ToLower(customer.SupplierEmailNumber))

// 	if customer.SupplierEmailNumber == "" {
// 		return utilerrors.NewBadRequestError("Invalid User Email")

// 	} else if !isEmailValid(customer.SupplierEmailNumber) {
// 		return utilerrors.NewBadRequestError("Invalid User Email")

// 	} else if customer.SupplierName == "" {
// 		return utilerrors.NewBadRequestError("Invalid User Name")
// 	}

// 	return nil
// }

func SelectStringElement(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}
