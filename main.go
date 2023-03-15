package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rarifbkn/api-rest-gorm-echo/database"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/api"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/repository"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/service"
	"github.com/rarifbkn/api-rest-gorm-echo/settings"
	"go.uber.org/fx"
)

// clean arquitechture
// ninguna capa tiene referencias a una capa superior... repository----> service----->presetacion(Api)----->
func main() {
	app := fx.New(
		// en provide van todas las funciones que devuelvan un struct
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		//en invoke van los comandos que quiero que se ejecuten antes de que corra la app
		fx.Invoke(
			setLifeCycle,
		),
	)
	app.Run()
}

// ciclo de vida de la app
func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{

		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
