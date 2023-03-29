package inventory

import (
	"context"
	"fmt"
	"log"
	"service/admin/case1/internal/inventory/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProviderMainService struct {
	client *mongo.Client
}

func (o ProviderMainService) RegisterProvider(currentUser string, providerChan <-chan models.ProviderInformation) {
	res := <-providerChan
	go func() {

		mongoCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		var con []models.ProviderContact
		var add []models.ProviderAddress
		var prod []models.ProviderProducts

		for c := range res.ProviderContacts {
			contact := models.ProviderContact{
				ContactName:          res.ProviderContacts[c].ContactName,
				ContactLastName:      res.ProviderContacts[c].ContactLastName,
				ContactPhoneNumber:   res.ProviderContacts[c].ContactPhoneNumber,
				ContactMessageNumber: res.ProviderContacts[c].ContactMessageNumber,
				ContactEmail:         res.ProviderContacts[c].ContactEmail,
			}
			con = append(con, contact)
		}
		for a := range res.ProviderAddress {
			address := models.ProviderAddress{
				ProviderStreet:  res.ProviderAddress[a].ProviderStreet,
				ProviderNumber:  res.ProviderAddress[a].ProviderNumber,
				ProviderZipCode: res.ProviderAddress[a].ProviderZipCode,
				ProviderCity:    res.ProviderAddress[a].ProviderCity,
			}
			add = append(add, address)
		}
		products := res.Validations(res)
		fmt.Println(products)
		for p := range res.ProviderProducts {
			products := models.ProviderProducts{
				ProductName:           res.ProviderProducts[p].ProductName,
				ProductCategory:       res.ProviderProducts[p].ProductCategory,
				ProductPrice:          res.ProviderProducts[p].ProductPrice,
				ProductMaterial:       res.ProviderProducts[p].ProductMaterial,
				ProductClassification: res.ProviderProducts[p].ProductClassification,
				ProductSKU:            products.ProviderProducts[p].ProductSKU,
			}
			prod = append(prod, products)
		}

		conn := o.client.Database("admin_db")
		coll := conn.Collection("admin_provider")

		providerInfo := models.ProviderInformation{
			ID:                primitive.NewObjectID(),
			ProviderBuissness: res.ProviderBuissness,
			ProviderContacts:  con,
			ProviderAddress:   add,
			ProviderProducts:  prod,
			CreatedAt:         time.Now().GoString(),
			AdminUser:         currentUser,
		}

		var err error
		_, err = coll.InsertOne(mongoCtx, providerInfo)
		if err != nil {
			log.Println("No User found ...", err.Error())
		}

	}()

}

func NewProviderService(client *mongo.Client) ProviderMainService {
	return ProviderMainService{client}
}
