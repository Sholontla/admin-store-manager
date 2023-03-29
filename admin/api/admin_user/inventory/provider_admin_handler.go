package inventory

import (
	"fmt"
	"log"
	"service/admin/case1/internal/inventory"
	"service/admin/case1/internal/inventory/models"
	middle "service/admin/case1/internal/middleware"
	loggs "service/admin/case1/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type ProviderHandlerService struct {
	Provider inventory.ProviderMainService
}

func (s ProviderHandlerService) ProviderRegister(ctxF *fiber.Ctx) error {

	logger, buffer := loggs.CreateLogger()
	var request models.ProviderInformation

	requestChan := make(chan models.ProviderInformation)

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println(fiber.StatusBadRequest, "Invalid JSON BDOY", err)
	}

	payLoad, err := middle.GetUserLogin(ctxF)
	if err != nil {
		return err
	}

	go func() {
		s.Provider.RegisterProvider(payLoad, requestChan)
	}()
	requestChan <- request
	close(requestChan)

	logger.Info("InventoryProviderRegister function running ...")
	log := loggs.Log{Encoded: buffer.String()}
	fmt.Println("Log before sending: ", log.Encoded)
	// loggs.GrcpClient(log.Encoded)
	return ctxF.JSON(fiber.Map{"admin_user": request})
}
