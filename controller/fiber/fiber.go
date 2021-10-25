package fiber

import (
	authUseCase "github.com/NV4RE/clarchgo/use-case/auth"
	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app    *fiber.App
	authUC *authUseCase.UseCase
}

func (f *fiberServer) Serve(addr string) error {
	return f.app.Listen(addr)
}

func NewFiberServer(authUC *authUseCase.UseCase) *fiberServer {
	app := fiber.New(fiber.Config{
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
	})

	f := &fiberServer{
		app:    app,
		authUC: authUC,
	}

	f.addRouteSystem()
	f.addRouteAuth()

	return f
}
