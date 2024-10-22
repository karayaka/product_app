package main

import (
	"context"
	"product_app/common/app"
	"product_app/common/postgresql"
	"product_app/controller"
	"product_app/persistence"
	"product_app/service"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	e := echo.New()

	configManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configManager.PostgreSqlConfig)

	pr := persistence.NewProductRepository(dbPool)

	ps := service.NemProductService(pr)

	pc := controller.NewProductController(ps)

	pc.RegisterRoutes(e)

	e.Start("localhost:8080")

}
