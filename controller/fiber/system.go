package fiber

import (
	"github.com/gofiber/fiber/v2"
)

func (f *fiberServer) addRouteSystem() {
	system := f.app.Group("/system")

	system.Get("/health", func(c *fiber.Ctx) error {
		//err = f.authRepo.Ok()
		//if err != nil {
		//	return f.errorHandler(c, err)
		//}

		return f.sendResponse(c, OK)
	})

}
