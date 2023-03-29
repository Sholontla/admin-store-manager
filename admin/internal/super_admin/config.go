package super_admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"service/admin/case1/internal/config"
	"service/admin/case1/internal/middleware"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SuperUserConfig struct {
	ID           uuid.UUID
	UserName     string
	UserPassword string
	CreatedAt    string
}

func compareCredentials(ctx *fiber.Ctx, access SuperAdmin, user string, password string) (bool, error) {

	if err := ComparePassword(password, access.Password); err != nil {
		return false, ctx.JSON(fiber.Map{"error": "Invalid_access"})
	}
	if user != access.UserName {
		return false, ctx.JSON(fiber.Map{"error": "Invalid_access"})
	}

	return true, nil
}

func SuperAdminReadDao(ctx *fiber.Ctx, requestChan <-chan SuperAdmin) string {

	request := <-requestChan
	u, p, _, _ := config.SuperUseConfig()
	t, err := compareCredentials(ctx, request, u, p)
	if err != nil {
		log.Println(err)
	}
	isSuperAdmin := strings.Contains(ctx.Path(), "/service/super/admin")
	fmt.Println(isSuperAdmin)
	if t {

		role := "super_admin"
		permissions := "all"

		var scope string
		if isSuperAdmin {
			scope = "super_admin"
		} else {
			scope = "admin"
		}
		nowTime := time.Now()
		expireTime := nowTime.Add(12 * time.Hour)

		token, errL := middleware.GenerateJWT(request.UserName, scope, []string{role}, []string{permissions})
		if errL != nil {
			ctx.Status(http.StatusOK).JSON(&fiber.Map{
				"message": "Invalid Credentials ..."})
		}

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  expireTime,
			HTTPOnly: true,
		}

		ctx.Cookie(&cookie)

		return "{access: successed}"

	} else {
		return "{error: invalid}"
	}
}

// fuction to activate read_write Only options

func SuperAdminRegisterDao(ctx *fiber.Ctx, requestChan <-chan SuperAdmin) {
	// fuction to activate read_write Only options

	go func() {
		for request := range requestChan {

			filea, _ := json.MarshalIndent(request, "", " ")
			os.WriteFile("read_write.json", filea, 0777) //0644
		}
	}()
}

func LogChannelProcess(s <-chan string) {

	go func() {
		logger := <-s
		st := map[string]string{
			"message": logger,
		}
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		LogPersonal(st["message"])

	}()
}

func LogPersonal(message string) {
	// open file and create if non-existent

	file, err := os.OpenFile("E:\\golang\\backend\\fintech\\projects\\case1\\admin\\internal\\super_admin\\admin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "message ", log.LstdFlags|log.Lshortfile)
	logger.Output(2, message)

}
