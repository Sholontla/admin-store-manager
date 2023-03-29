package inventory

import (
	"service/admin/case1/internal/db"
	"service/admin/case1/internal/inventory"
	security "service/admin/case1/internal/middleware"
	"service/admin/case1/pkg/monitoring/middleware"

	"github.com/gofiber/fiber/v2"
)

type RolesPermissions struct {
	PermissionsRegister []string
	PermissionsUpdate   []string
	ReadRegister        []string
	Roles               []string
}

var (
	roles           = []string{"admin"}
	permInvRegister = []string{"create_inventory"}
	permInvUpdate   = []string{"update_inventory"}
)

func AdminUsersRouters(app *fiber.App) {
	rp := RolesPermissions{
		PermissionsRegister: permInvRegister,
		PermissionsUpdate:   permInvUpdate,
		Roles:               roles,
	}

	serviceInv := inventory.NewInventoryService(db.ConnMongoDB())
	serviceProv := inventory.NewProviderService(db.ConnMongoDB())
	inv := InventoryHandlerService{serviceInv}
	prov := ProviderHandlerService{serviceProv}

	s := app.Group("service")
	a := s.Group("admin")
	p := a.Group("provider", middleware.RecordRequestLatency)

	adminAuthenticated := p.Use(security.IsAuthenticated)

	permissionRegister := adminAuthenticated.Use(security.ValidateRolesAndPermissions(rp.Roles, rp.PermissionsRegister))

	permissionRegister.Post("register", prov.ProviderRegister)
	permissionRegister.Post("get/inventory/:provider", inv.InventoryProviderRegister)
	permissionRegister.Get("get/all/inventory", inv.GetAllProvidersHandler)
	permissionRegister.Post("register/order/:provider", inv.CreateInventoryOrder)
	permissionRegister.Post("filter/providers/:providers", inv.FilterByProvidersHandler)
}
