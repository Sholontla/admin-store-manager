package server

import (
	"log"
	"os"
	"os/signal"
	"service/admin/case1/api/admin_user/admin"
	inventory "service/admin/case1/api/admin_user/inventory"
	superadmin "service/admin/case1/api/super_admin"
	"service/admin/case1/internal/messaging/grpc"
	"service/admin/case1/pkg/monitoring"
	"service/admin/case1/pkg/monitoring/middleware"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func AdminUserServerStart() {

	app := fiber.New()
	superadmin.SuperAdminUsersRouters(app)
	admin.AdminUsersRouters(app)
	inventory.AdminUsersRouters(app)
	monitoring.PrometheusRoute(app)
	middleware.RegisterPrometheusMetrics()

	go app.Listen(":8001")

	grpc.GrpcProductServer()
	// SIGINT is the signal sent when we press Ctrl+C
	// SIGTERM gracefully kills the process
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down Service Demo server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Service Demo Server: %v\n", err)
	}
}
