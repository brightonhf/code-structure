package ports

import "code_structure/internals/domain/models"

type MenuService interface {
	GetMeals() models.Meal
}
