package model

type MenuType string

type MenuItem struct {
	OrderCode string `gorm:"primaryKey"`
	Name      string
	Price     float64
	Type      MenuType
}
