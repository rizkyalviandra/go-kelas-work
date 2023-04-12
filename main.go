package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     float64
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Fried Rice",
			OrderCode: "F-0000001",
			Price:     70000,
		},
		{
			Name:      "Fried Fries",
			OrderCode: "F-0000002",
			Price:     30000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}
func getDrinkMenu(c echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name:      "Orange Juice",
			OrderCode: "F-0000001",
			Price:     20000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()

	e.GET("/foods", getFoodMenu)
	e.GET("/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":8000"))
}
