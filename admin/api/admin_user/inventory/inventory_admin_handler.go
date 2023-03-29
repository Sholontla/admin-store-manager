package inventory

import (
	"fmt"
	"log"
	"service/admin/case1/internal/inventory"
	middle "service/admin/case1/internal/middleware"
	"service/admin/case1/internal/super_admin"
	loggs "service/admin/case1/pkg/logger"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type InventoryHandlerService struct {
	Inventory inventory.InventoryMainService
}

var wg sync.WaitGroup

func (InventoryHandlerService) InventoryProviderRegister(ctxF *fiber.Ctx) error {
	logger, buffer := loggs.CreateLogger()

	provider := ctxF.Params("provider")

	prov := <-inventory.FindByProduct(provider)

	logChan := make(chan string)

	logger.Info("AdminRegister function running ...")
	log := loggs.Log{Encoded: buffer.String()}
	l := fmt.Sprintf("Log before sending: %s", log.Encoded)
	wg.Add(1)
	go func() {
		defer wg.Done()
		super_admin.LogChannelProcess(logChan)
	}()

	logChan <- l
	close(logChan)
	wg.Wait()

	return ctxF.JSON(fiber.Map{"provider": prov})
}

type OrderParams struct {
	ProductName     []string `json:"product_name"`
	ProductQuantity []int    `json:"product_quantity"`
}

func (i InventoryHandlerService) GetAllProvidersHandler(ctx *fiber.Ctx) error {
	res, err := i.Inventory.GetAllProvidersDAO()
	if err != nil {
		ctx.JSON(fiber.Map{"error": err.Error()})
		return err
	}
	return ctx.JSON(fiber.Map{"provider": res})
}

func (s InventoryHandlerService) CreateInventoryOrder(ctx *fiber.Ctx) error {
	var request OrderParams

	if err := ctx.BodyParser(&request); err != nil {
		log.Println(err)
	}

	payLoad, err := middle.GetUserLogin(ctx)
	if err != nil {
		return err
	}

	logger, buffer := loggs.CreateLogger()

	prov := ctx.Params("provider")

	productNames := request.ProductName
	productQuantity := request.ProductQuantity
	fmt.Println(productNames, productQuantity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Inventory.RegisterInventoryOrder(ctx, payLoad, prov, productNames, productQuantity)
	}()
	wg.Wait()

	logger.Info("InventoryProviderRegister function running ...")
	log := loggs.Log{Encoded: buffer.String()}
	fmt.Println("Log before sending: ", log.Encoded)
	// loggs.GrcpClient(log.Encoded)
	return ctx.JSON(fiber.Map{"admin_user": "request"})
}

func (InventoryHandlerService) FilterByProvidersHandler(ctx *fiber.Ctx) error {
	providers := ctx.Query("providers")
	providerList := strings.Split(providers, "/")
	results := make([]inventory.InventoryOrder, 0)

	var request map[string]int

	if err := ctx.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	for _, provider := range providerList {
		providerChan := inventory.FilterByProviders(ctx, request, []string{provider})
		result := <-providerChan
		results = append(results, result)
		// 	if result.ProviderBuissness.ProviderBuissness != "" {
		//
		// 	}
	}

	return ctx.JSON(results)
}
