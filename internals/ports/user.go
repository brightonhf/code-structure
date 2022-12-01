package ports

import (
	"code_structure/internals/domain/models"
)

type FetchUseCase interface {
	FetchByID(ID string) models.User
	FetchAll() []models.User
}

type UpdateUseCase interface {
	UpdateByID(ID string) models.User
}

// Secondary ports

// FetchRepository Secondary Port
type FetchRepository interface {
}

// UpdateRepository secondary Port
type UpdateRepository interface {
}
