package menu

import "github.com/rizkyalviandra/go-kelas-work/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
