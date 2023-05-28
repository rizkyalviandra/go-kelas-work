package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rizkyalviandra/go-kelas-work/internal/database"
	"github.com/rizkyalviandra/go-kelas-work/internal/delivery/rest"
	"github.com/rizkyalviandra/go-kelas-work/internal/repository/menu"
	"github.com/rizkyalviandra/go-kelas-work/internal/usecase/resto"
)

const (
	dsn = "host=localhost user=postgres password=postgres dbname=go_rest_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func main() {
	e := echo.New()

	db := database.GetDB(dsn)

	menuRepo := menu.GetRepository(db)
	restoUsecase := resto.GetUsecase(menuRepo)
	h := rest.NewHandler(restoUsecase)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":8000"))
}
