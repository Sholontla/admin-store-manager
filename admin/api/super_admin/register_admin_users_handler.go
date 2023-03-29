package super_admin

import (
	"fmt"
	"log"
	"net/http"
	"service/admin/case1/internal/models"
	"service/admin/case1/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterAdminUserHandler(ctxF *fiber.Ctx) error {

	var request models.AdminUsers
	requestChan := make(chan models.AdminUsers)

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	var rl []string

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := services.RegisterAdminUser(ctxF.Context(), requestChan)
		if err != nil {
			rl = append(rl, fmt.Sprintf("Error: %s", err.Error()))
			return
		}
		rl = append(rl, "register")
	}()

	requestChan <- request
	close(requestChan)
	wg.Wait()

	if len(rl) > 0 {
		return ctxF.Status(500).JSON(fiber.Map{"error": rl[0]})
	}

	return ctxF.JSON(fiber.Map{"admin_user": rl})
}

func UpdateAdminUserHandler(c *fiber.Ctx) error {

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

// func DeleteAdminUserHanlder(ctxF *fiber.Ctx) error {

// 	logger, buffer := loggs.CreateLogger()

// 	var request super_admin.SuperAdmin

// 	if err := ctxF.BodyParser(&request); err != nil {
// 		log.Println("Invalid JSON BDOY", err)
// 	}
// 	if len(request.UserName) == 0 {
// 		return ctxF.JSON(fiber.Map{"user_email": "User Email is required ..."})
// 	}
// 	if len(request.Password) < 7 {
// 		return ctxF.JSON(fiber.Map{"password": "Password must be at least lenght of 8 elements ..."})
// 	}

// 	requestChan := make(chan super_admin.SuperAdmin)
// 	permissionsChan := make(chan super_admin.SuperAdmin)

// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		super_admin.AdminRegisterDao(ctxF, requestChan)
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		super_admin.AdminPermissionsDao(ctxF, permissionsChan)
// 	}()

// 	requestChan <- request
// 	permissionsChan <- request
// 	close(requestChan)
// 	close(permissionsChan)
// 	wg.Wait()

// 	logger.Info("AdminRegister function running ...")
// 	log := loggs.Log{Encoded: buffer.String()}
// 	fmt.Println("Log before sending: ", log.Encoded)
// 	// loggs.GrcpClient(log.Encoded)
// 	return ctxF.JSON(fiber.Map{"admin_user": request})
// }
