package fiber

import (
	"errors"
	"github.com/NV4RE/clarchgo/entity/auth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var (
	ErrAuthorizationHeaderRequired = errors.New("authorization header required")
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

	ah := strings.Split(authHeader, " ")
	if len(ah) != 2 {
		return auth.User{}, ErrAuthorizationHeaderRequired
	}

	return f.authUC.GetUserByToken(c.Context(), ah[1])
}
