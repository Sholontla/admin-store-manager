package super_admin

import (
	"fmt"
	"log"
	"net/http"
	"service/admin/case1/internal/super_admin"
	loggs "service/admin/case1/pkg/logger"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup

func SuperAdminAccess(ctxF *fiber.Ctx) error {
	logger, buffer := loggs.CreateLogger()

	var request super_admin.SuperAdmin

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	requestPassChan := make(chan super_admin.SuperAdmin)
	var rl []string
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := super_admin.SuperAdminReadDao(ctxF, requestPassChan)
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

func SuperAdminLogout(c *fiber.Ctx) error {

	nowTime := time.Now()
	expireTime := nowTime.Add(-time.Hour)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  expireTime,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success Logout"})

}

func SuperAdminUpdate(ctxF *fiber.Ctx) error {

	logger, buffer := loggs.CreateLogger()

	var request super_admin.SuperAdmin

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}
	if len(request.UserName) == 0 {
		return ctxF.JSON(fiber.Map{"user_email": "User Email is required ..."})
	}
	if len(request.Password) < 7 {
		return ctxF.JSON(fiber.Map{"password": "Password must be at least lenght of 8 elements ..."})
	}

	requestChan := make(chan super_admin.SuperAdmin)
	permissionsChan := make(chan super_admin.SuperAdmin)

	wg.Add(1)
	go func() {
		defer wg.Done()
		super_admin.AdminRegisterDao(ctxF, requestChan)
	}()

	requestChan <- request
	permissionsChan <- request
	close(requestChan)
	close(permissionsChan)
	wg.Wait()

	logger.Info("AdminRegister function running ...")
	log := loggs.Log{Encoded: buffer.String()}
	fmt.Println("Log before sending: ", log.Encoded)
	// loggs.GrcpClient(log.Encoded)
	return ctxF.JSON(fiber.Map{"admin_user": request})
}

// func SuperAdminDelete(ctxF *fiber.Ctx) error {

// 	logger, buffer := loggs.CreateLogger()

// 	var request models.AdminUsers

// 	if err := ctxF.BodyParser(&request); err != nil {
// 		log.Println("Invalid JSON BDOY", err)
// 	}
// 	if len(request.AdminUserEmail) == 0 {
// 		return ctxF.JSON(fiber.Map{"user_email": "User Email is required ..."})
// 	}
// 	if len(request.Password) < 7 {
// 		return ctxF.JSON(fiber.Map{"password": "Password must be at least lenght of 8 elements ..."})
// 	}

// 	perm := request

// 	requestChan := make(chan models.AdminUsers)
// 	permissionsChan := make(chan models.AdminUsers)

// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		services.AdminRegisterDao(ctxF, requestChan)
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		services.AdminPermissionsDao(ctxF, permissionsChan)
// 	}()

// 	requestChan <- request
// 	permissionsChan <- perm
// 	close(requestChan)
// 	close(permissionsChan)
// 	wg.Wait()

// 	logger.Info("AdminRegister function running ...")
// 	log := loggs.Log{Encoded: buffer.String()}
// 	fmt.Println("Log before sending: ", log.Encoded)
// 	// loggs.GrcpClient(log.Encoded)
// 	return ctxF.JSON(fiber.Map{"admin_user": request})
// }
