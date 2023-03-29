package admin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"service/admin/case1/internal/db"
	"service/admin/case1/internal/middleware"
	"service/admin/case1/internal/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func compareCredentials(ctx *fiber.Ctx, access models.AdminUserLogin, user string, password string) (bool, error) {

	if err := models.ComparePassword(password, access.Password); err != nil {
		return false, ctx.JSON(fiber.Map{"error": "Invalid_access"})
	}
	if user != access.AdminUserEmail {
		return false, ctx.JSON(fiber.Map{"error": "Invalid_access"})
	}

	return true, nil
}

func AdminUserLogin(ctx *fiber.Ctx, requestChan <-chan models.AdminUserLogin) string {

	request := <-requestChan
	findChan := FindAdminUserByEmail(request.AdminUserEmail)
	user := <-findChan

	t, err := compareCredentials(ctx, request, user.AdminUserEmail, user.Password)
	if err != nil {
		log.Println(err)
	}

	isSuperAdmin := strings.Contains(ctx.Path(), "/service/admin")
	var scope string
	if isSuperAdmin {
		scope = "admin"
	} else {
		scope = "super_admin"
	}
	per := RolesPermissionsAssignment(user)

	if t {

		nowTime := time.Now()
		expireTime := nowTime.Add(12 * time.Hour)
		fmt.Println(per)
		token, errL := middleware.GenerateJWT(user.AdminUserEmail, scope, []string{"admin"}, per)
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

func FindAdminUserByEmail(email string) chan models.AdminUsers {
	adminChan := make(chan models.AdminUsers)
	go func() {
		var adminUser models.AdminUsers
		condition := bson.M{
			"admin_user_email": email,
		}
		mongoCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		conn := db.MongoConn.Database("admin_db")
		coll := conn.Collection("admin_users")
		err := coll.FindOne(mongoCtx, condition).Decode(&adminUser)
		if err != nil {
			log.Println("No User found ...", err.Error())
			adminChan <- models.AdminUsers{} // send empty model to channel
			return
		}

		adminChan <- adminUser // send model to channel
	}()

	return adminChan
}
