package main

import (
	"net/http"
	"os"

	"github.com/LucasGabrielBravo/primo-block-store/controllers"
	"github.com/LucasGabrielBravo/primo-block-store/usecases"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	loadPublic(app)
	loadRoutes(app)

	err := app.Start()

	if err != nil {
		panic(err)
	}
}

func loadPublic(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		return nil
	})
}

func loadRoutes(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		g := e.Router.Group("/api/v1")

		g.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/status",
			Handler: func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]any{"status": "ok"})
			},
		})

		g.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/symbollist/:id",
			Handler: controllers.NewGetSymbolList(usecases.NewGetSymbolList(app)).Handler,
		})

		return nil
	})
}
