package admin

import (
	"fmt"
	"log"
	"service/admin/case1/internal/models"
	"service/admin/case1/internal/services/admin"
	"service/admin/case1/internal/super_admin"
	loggs "service/admin/case1/pkg/logger"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup

func AdminUserLogin(ctxF *fiber.Ctx) error {
	logger, buffer := loggs.CreateLogger()

	var request models.AdminUserLogin

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	requestPassChan := make(chan models.AdminUserLogin)
	var rl []string
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := admin.AdminUserLogin(ctxF, requestPassChan)
		rl = append(rl, r)
	}()

	requestPassChan <- request
	close(requestPassChan)
	wg.Wait()

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

	return ctxF.JSON(fiber.Map{"admin_user": rl})
}
