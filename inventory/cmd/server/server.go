package server

import (
	"log"
	"os"
	"os/signal"
	"service/inventory/case1/pkg/monitoring"
	"service/inventory/case1/pkg/monitoring/middleware"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func ProducerServerStart() {

	app := fiber.New()

	monitoring.PrometheusRoute(app)
	middleware.RegisterPrometheusMetrics()

	go app.Listen(":8001")

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
