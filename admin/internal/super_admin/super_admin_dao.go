package super_admin

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
)

func AdminRegisterDao(ctx *fiber.Ctx, requestChan <-chan SuperAdmin) {
	// fuction to activate read_write Only options

	go func() {
		for request := range requestChan {

			filea, _ := json.MarshalIndent(request, "", " ")
			os.WriteFile("read_write.json", filea, 0777) //0644
		}
	}()
}
