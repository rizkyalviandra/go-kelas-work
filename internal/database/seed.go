package database

import (
	"github.com/rizkyalviandra/go-kelas-work/internal/model"
	"github.com/rizkyalviandra/go-kelas-work/internal/model/constants"
	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			OrderCode: "F-0000001",
			Name:      "Fried Rice",
			Price:     70000,
			Type:      constants.MenuTypeFood,
		},
		{
			OrderCode: "F-0000002",
			Name:      "Fried Fries",
			Price:     30000,
			Type:      constants.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			OrderCode: "D-0000001",
			Name:      "Orange Juice",
			Price:     20000,
			Type:      constants.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
