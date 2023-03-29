package order

import (
	"context"
	"log"
	"reflect"
	inv "service/stores/case1/internal/inventory/models"
	"service/stores/case1/internal/models"
	"service/stores/case1/pkg/concurrency"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dataBase   = "stores"
	collection = "stores-a"
)

type OrderMainService struct {
	client *mongo.Client
}

func (o OrderMainService) CreateOrderDAO(orderRequestChan <-chan models.CustomerOrderCompleted) Ticket {

	orderPool, err := concurrency.GenericPool(reflect.TypeOf(models.CustomerOrderCompleted{}))
	if err != nil {
		log.Println(err)
	}

	orderCreated := orderPool.Get().(chan models.CustomerOrderCompleted)
	defer orderPool.Put(orderCreated)

	ticketChan := make(chan Ticket)
	orderRequest := <-orderRequestChan

	done := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var allProducts []inv.Product
	for _, product := range orderRequest.Inventory.InventoryProduct {
		allProducts = append(allProducts, inv.Product{
			ProductTitle:        product.ProductTitle,
			ProductModel:        product.ProductModel,
			ProductMaterial:     product.ProductMaterial,
			ProductPrice:        product.ProductPrice,
			ProductQuantity:     product.ProductQuantity,
			ProductSerialNumber: product.ProductSerialNumber,
		})
	}
	a := Finance{
		ID: orderRequest.ID,
		Inventory: inv.Inventory{
			ID:               orderRequest.Inventory.ID,
			InventoryProduct: allProducts,
			CreatedAt:        orderRequest.Inventory.CreatedAt,
		},
		StoreInfomation: models.StoreInfomation{},
		CreatedAt:       orderRequest.CreatedAt,
	}

	b := Accounting{
		ID: orderRequest.ID,
		Inventory: inv.Inventory{
			ID:               orderRequest.Inventory.ID,
			InventoryProduct: allProducts,
			CreatedAt:        orderRequest.Inventory.CreatedAt,
		},
		StoreInfomation: models.StoreInfomation{},
		CreatedAt:       orderRequest.CreatedAt,
	}

	c := Admin{
		ID: orderRequest.ID,
		Inventory: inv.Inventory{
			ID:               orderRequest.Inventory.ID,
			InventoryProduct: allProducts,
			CreatedAt:        orderRequest.Inventory.CreatedAt,
		},
		StoreInfomation: models.StoreInfomation{},
		CreatedAt:       orderRequest.CreatedAt,
	}

	e := &Marketing{
		ID: orderRequest.ID,
		Inventory: inv.Inventory{
			ID:               orderRequest.Inventory.ID,
			InventoryProduct: allProducts,
			CreatedAt:        orderRequest.Inventory.CreatedAt,
		},
		StoreInfomation: models.StoreInfomation{},
		Employee:        models.EmployeeInternalInformation{},
		CreatedAt:       orderRequest.CreatedAt,
	}

	GetProcessor(a, b, c, *e)
	orderRequest.TicketHash = uuid.New()
	conn := o.client.Database(dataBase)
	col := conn.Collection(collection)
	_, errInsert := col.InsertOne(ctx, orderRequest)
	if errInsert != nil {
		log.Println(errInsert)
	}
	go func() {
		t := orderCompletedDAO(orderCreated)
		close(done)
		ticketChan <- t
	}()

	orderCreated <- orderRequest

	<-done
	t := <-ticketChan
	close(ticketChan)

	return t
}

func orderCompletedDAO(orderCompletedChan <-chan models.CustomerOrderCompleted) Ticket {
	d := <-orderCompletedChan
	coc := createTicket(d)
	s := <-coc
	return s
}

func createTicket(createTicket models.CustomerOrderCompleted) chan Ticket {
	coc := make(chan Ticket)
	done := make(chan struct{})
	go func() {
		defer close(done)
		var allProducts []TicketProducts
		for _, product := range createTicket.Inventory.InventoryProduct {
			allProducts = append(allProducts, TicketProducts{
				ProductTitle:        product.ProductTitle,
				ProductModel:        product.ProductModel,
				ProductMaterial:     product.ProductMaterial,
				ProductPrice:        product.ProductPrice,
				ProductQuantity:     product.ProductQuantity,
				ProductSerialNumber: product.ProductSerialNumber,
			})
		}

		// create ticket process
		ticket := Ticket{
			ID:             primitive.NewObjectID(),
			OrderReference: "",
			Products:       allProducts,
			StoreInfomation: StoreInfomation{
				StoreName:        createTicket.StoreInfomation.StoreName,
				StoreCode:        createTicket.StoreInfomation.StoreCode,
				StorePhoneNumber: createTicket.StoreInfomation.StorePhoneNumber,
				StoreEmail:       createTicket.StoreInfomation.StoreEmail,
			},
			Employee: EmployeeInternalInformation{
				Area:        createTicket.Employee.Area,
				Designation: createTicket.Employee.Designation,
			},
			CreatedAt:  time.Now().Local().String(),
			TicketHash: createTicket.TicketHash,
		}

		coc <- ticket

		close(coc)
	}()

	return coc
}

type Ticket struct {
	ID              primitive.ObjectID
	OrderReference  string
	Products        []TicketProducts
	StoreInfomation StoreInfomation
	Employee        EmployeeInternalInformation
	CreatedAt       string
	TicketHash      uuid.UUID
}

type TicketProducts struct {
	ProductTitle        string
	ProductModel        string
	ProductMaterial     string
	ProductPrice        float64
	ProductQuantity     int64
	ProductSerialNumber string
}

type StoreInfomation struct {
	StoreName        string
	StoreCode        string
	StorePhoneNumber string
	StoreEmail       string
}

type EmployeeInternalInformation struct {
	Area        string
	Designation string
}

func NewOrderTicket(client *mongo.Client) OrderMainService {
	return OrderMainService{client}
}
