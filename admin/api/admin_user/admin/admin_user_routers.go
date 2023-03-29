package admin

import (
	"service/admin/case1/pkg/monitoring/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminUsersRouters(app *fiber.App) {

	s := app.Group("service")
	p := s.Group("admin", middleware.RecordRequestLatency)

	p.Post("login", AdminUserLogin)

	// adminAuthenticated := p.Use(midleware.IsAuthenticated)

	// adminAuthenticated.Post("update", SuperAdminUpdate)
	// adminAuthenticated.Post("logout", SuperAdminLogout)
	// adminAuthenticated.Post("register/admin", RegisterAdminUserHandler)
	// p.Delete("delete", SuperAdminDelete)
}
