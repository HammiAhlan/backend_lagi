package url

import (
	"github.com/HammiAhlan/backend_lagi/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	
	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik
	
	page.Get("/donasi", controller.GetDonasiData)
}

