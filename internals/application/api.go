package application

import (
	"code_structure/internals/domain/selection"
	"code_structure/internals/ports"
)

type restAPI struct {
	menuService ports.MenuService
	core        selection.Core
}

func NewAPI(menuAdapter ports.MenuService, core selection.Core) ports.API {
	return &restAPI{menuService: menuAdapter, core: core}
}

func (r restAPI) GetSelection() (int, error) {

	meal := r.menuService.GetMeals()

	res := r.core.CountSelection(len(meal.Name))

	return res, nil
}
