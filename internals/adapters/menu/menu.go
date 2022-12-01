package menuadapter

import (
	"net/http"

	"code_structure/internals/domain/models"
	"code_structure/internals/ports"
)

type menuAdapter struct {
	client *http.Client
}

func NewMenuAdapter(client *http.Client) ports.MenuService {
	return &menuAdapter{client: client}
}

func (m menuAdapter) GetMeals() models.Meal {
	return models.Meal{}
}
