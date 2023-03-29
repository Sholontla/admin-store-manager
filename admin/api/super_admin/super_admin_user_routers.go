package super_admin

import (
	midleware "service/admin/case1/internal/middleware"
	"service/admin/case1/pkg/monitoring/middleware"

	"github.com/gofiber/fiber/v2"
)

func SuperAdminUsersRouters(app *fiber.App) {

	s := app.Group("service")
	p := s.Group("super/admin", middleware.RecordRequestLatency)

	p.Post("login", SuperAdminAccess)

	adminAuthenticated := p.Use(midleware.IsAuthenticated)

	adminAuthenticated.Post("update", SuperAdminUpdate)
	adminAuthenticated.Post("logout", SuperAdminLogout)
	adminAuthenticated.Post("register/admin", RegisterAdminUserHandler)
	// p.Delete("delete", SuperAdminDelete)
}
