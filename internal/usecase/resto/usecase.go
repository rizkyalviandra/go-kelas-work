package resto

import "github.com/rizkyalviandra/go-kelas-work/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
