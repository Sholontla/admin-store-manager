package server

import (
	"log"
	"os"
	"os/signal"
	"service/stores/case1/api/admin"
	"service/stores/case1/api/stores"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func ServerStart() {

	app := fiber.New()

	stores.StoreRouters(app)
	admin.AdminRouters(app)

	go app.Listen(":8002")

	// SIGINT is the signal sent when we press Ctrl+C
	// SIGTERM gracefully kills the process
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Banking/Transaction Server: %v\n", err)
	}
}
