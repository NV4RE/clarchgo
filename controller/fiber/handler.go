package fiber

import (
	"clarchgo/entity/auth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

const OK string = "OK"

type Response struct {
	Data      interface{} `json:"data"`
	Error     string      `json:"error"`
	ErrorCode int         `json:"error_code"`
}

func (f *fiberServer) sendResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(&Response{
		Data: data,
	})
}

func (f *fiberServer) sendError(c *fiber.Ctx, status int, err string, errCode int) error {
	return c.Status(status).JSON(&Response{
		Error:     err,
		ErrorCode: errCode,
	})
}

func (f *fiberServer) errorHandler(c *fiber.Ctx, err error) error {
	//switch err {
	//case auth_repo.ErrNotFound:
	//	return f.sendError(c, 404, err.Error(), 1)
	//}

	return f.sendError(c, 500, err.Error(), 0)
}

func (f *fiberServer) getUserFromHeader(c *fiber.Ctx) (auth.User, error) {
	authHeader := string(c.Request().Header.Peek("Authorization"))

	ah := strings.SplitN(authHeader, " ", 2)
	return f.authUC.GetUserByToken(c.Context(), ah[1])
}
