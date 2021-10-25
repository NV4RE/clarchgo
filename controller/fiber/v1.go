package fiber

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type createUserWithPasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (f fiberServer) addRouteAuth() {
	a := f.app.Group("/auth")
	a.Use(logger.New())

	a.Post("/", func(c *fiber.Ctx) error {
		usr, err := f.getUserFromHeader(c)
		if err != nil {
			return f.errorHandler(c, err)
		}

		var upr createUserWithPasswordRequest
		err = json.Unmarshal(c.Body(), &upr)
		if err != nil {
			return f.errorHandler(c, err)
		}

		u, err := f.authUC.CreateUserWithPassword(c.Context(), usr, upr.Username, upr.Password)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return f.sendResponse(c, u)
	})

	a.Get("/:username", func(c *fiber.Ctx) error {
		usr, err := f.getUserFromHeader(c)
		if err != nil {
			return f.errorHandler(c, err)
		}

		u, err := f.authUC.GetUserByUsername(c.Context(), usr, c.Params("username"))
		if err != nil {
			return f.errorHandler(c, err)
		}

		return f.sendResponse(c, u)
	})

}
