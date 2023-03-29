package stores

import (
	"log"
	db "service/stores/case1/datasource/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type CashierOrderMainService struct {
	client *mongo.Client
}

func (o CashierOrderMainService) CashierOrderCompletedDAO(orderRequestChan <-chan map[string]interface{}) (chan map[string]interface{}, error) {

	chanOrderRequest := make(chan map[string]interface{})
	defer close(chanOrderRequest)

	orderRequest := <-orderRequestChan

	db.InsertOne(o.client, orderRequest)

	chanOrderRequest <- orderRequest

	return chanOrderRequest, nil
}

func (o CashierOrderMainService) GetAllInventoryDAO() ([]map[string]interface{}, error) {

	res, err := db.GetAll(o.client)
	if err != nil {
		log.Println(err)
	}
	return res, nil
}

func NewCashierOrderCompeleted(client *mongo.Client) CashierOrderMainService {
	return CashierOrderMainService{client}
}
