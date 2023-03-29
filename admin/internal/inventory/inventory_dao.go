package inventory

import (
	"context"
	"fmt"
	"log"
	"service/admin/case1/internal/db"
	"service/admin/case1/internal/inventory/models"
	middle "service/admin/case1/internal/middleware"
	"service/admin/case1/pkg/pdf"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collInventory = "inventory"
	collProvider  = "admin_provider"
	collOrderInv  = "admin_inventory_orders"
)

type InventoryMainService struct {
	client *mongo.Client
}

func FindByProduct(provider string) chan models.ProviderInformation {
	adminChan := make(chan models.ProviderInformation)
	go func() {
		var adminUser models.ProviderInformation
		condition := bson.M{
			"provider_buissness": provider,
		}
		mongoCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		conn := db.MongoConn.Database("admin_db")
		coll := conn.Collection("admin_provider")
		err := coll.FindOne(mongoCtx, condition).Decode(&adminUser)
		if err != nil {
			log.Println("No User found ...", err.Error())
			adminChan <- models.ProviderInformation{} // send empty model to channel
			return
		}

		adminChan <- adminUser // send model to channel
	}()

	return adminChan
}

func FilterByProviders(ctx *fiber.Ctx, request map[string]int, providers []string) chan InventoryOrder {
	adminChan := make(chan InventoryOrder)

	payLoad, err := middle.GetUserLogin(ctx)
	if err != nil {
		log.Println(err)
	}

	go func() {
		var adminUser models.ProviderInformation

		condition := bson.M{
			"provider_buissness": bson.M{"$in": providers},
		}
		mongoCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		conn := db.MongoConn.Database("admin_db")
		coll := conn.Collection("admin_provider")

		err := coll.FindOne(mongoCtx, condition).Decode(&adminUser)
		if err != nil {
			log.Println("No User found ...", err.Error())
			adminChan <- InventoryOrder{} // send empty model to channel
			return
		}

		var products []OrderProviderProducts

		for p := range adminUser.ProviderProducts {
			pr := OrderProviderProducts{
				ProductName:           adminUser.ProviderProducts[p].ProductName,
				ProductCategory:       adminUser.ProviderProducts[p].ProductCategory,
				ProductPrice:          adminUser.ProviderProducts[p].ProductPrice,
				ProductSKU:            adminUser.ProviderProducts[p].ProductSKU,
				ProductMaterial:       adminUser.ProviderProducts[p].ProductMaterial,
				ProductClassification: adminUser.ProviderProducts[p].ProductClassification,
				ProductQuantity:       int(request["quantity"]),
			}
			products = append(products, pr)
		}

		invOrderInfo := InventoryOrder{
			ID:                primitive.NewObjectID(),
			ProviderBuissness: adminUser.ProviderBuissness,
			ProviderContact:   adminUser.ProviderContacts[0],
			ProviderAddress:   adminUser.ProviderAddress[0],
			ProviderProduct:   products,
			CreatedAt:         time.Now().GoString(),
			UserAdmin:         payLoad,
			//OrderPayment: ,
		}

		adminChan <- invOrderInfo // send model to channel
	}()

	return adminChan
}

func (c InventoryMainService) RegisterInventoryOrder(ctx *fiber.Ctx, currentUser string, provider string, productNames []string, productQuantity []int) {

	go func() {
		prov := <-FindByProduct(provider)
		var filteredProviders []models.ProviderProducts
		for _, product := range prov.ProviderProducts {
			for i, name := range productNames {
				if product.ProductName == name {
					filteredProviders = append(filteredProviders, models.ProviderProducts{
						ProductName:           product.ProductName,
						ProductCategory:       product.ProductCategory,
						ProductPrice:          product.ProductPrice,
						ProductSKU:            product.ProductSKU,
						ProductMaterial:       product.ProductMaterial,
						ProductClassification: product.ProductClassification,
						ProductQuantity:       productQuantity[i],
					})
					break
				}
			}
		}

		var orderProducts []OrderProviderProducts
		for _, product := range filteredProviders {
			orderProducts = append(orderProducts, OrderProviderProducts(product))
		}

		invOrderInfo := InventoryOrder{
			ID:                primitive.NewObjectID(),
			ProviderBuissness: prov.ProviderBuissness,
			ProviderContact:   prov.ProviderContacts[0],
			ProviderAddress:   prov.ProviderAddress[0],
			ProviderProduct:   orderProducts,
			CreatedAt:         time.Now().GoString(),
			UserAdmin:         currentUser,
			//OrderPayment: ,
		}

		var products []models.ProviderProducts

		for pr := range invOrderInfo.ProviderProduct {
			products = append(products, models.ProviderProducts{
				ProductName:           invOrderInfo.ProviderProduct[pr].ProductName,
				ProductCategory:       invOrderInfo.ProviderProduct[pr].ProductCategory,
				ProductPrice:          invOrderInfo.ProviderProduct[pr].ProductPrice,
				ProductSKU:            invOrderInfo.ProviderProduct[pr].ProductSKU,
				ProductMaterial:       invOrderInfo.ProviderProduct[pr].ProductMaterial,
				ProductClassification: invOrderInfo.ProviderProduct[pr].ProductClassification,
				ProductQuantity:       invOrderInfo.ProviderProduct[pr].ProductQuantity,
			})

		}

		invInfo := InventoryInformation{
			ID:               primitive.NewObjectID(),
			ProviderBusniess: invOrderInfo.ProviderBuissness,
			ProviderProducts: products,
			CreatedAt:        time.Now().Local().GoString(),
			UserAdmin:        currentUser,
		}

		for _, product := range prov.ProviderProducts {
			for i := range productNames {
				if product.ProductName == productNames[i] {
					db.UpdateMany(c.client, provider, product.ProductSKU, productQuantity, collInventory)
				} else {
					db.InsertOne(c.client, invInfo, collInventory)
				}

				filteredProviders = append(filteredProviders, models.ProviderProducts{
					ProductSKU: product.ProductSKU,
				})
			}

		}

		db.InsertOne(c.client, invOrderInfo, collOrderInv)

		pdf.InventoryOrderPDF(Test(invOrderInfo), "invOrderInfo")

	}()

}

func (o InventoryMainService) GetAllProvidersDAO() ([]models.ProviderInformation, error) {

	res, err := db.GetAll(o.client, collProvider)
	if err != nil {
		log.Println(err)
	}
	return res, nil
}

func Test(inv InventoryOrder) [][]string {

	values := make([][]string, len(inv.ProviderProduct))

	for i, product := range inv.ProviderProduct {
		values[i] = []string{
			product.ProductName,
			product.ProductCategory,
			fmt.Sprintf("%f", product.ProductPrice),
			product.ProductSKU,
			product.ProductMaterial,
			product.ProductClassification,
		}
	}

	return values
}

func NewInventoryService(client *mongo.Client) InventoryMainService {
	return InventoryMainService{client}
}
