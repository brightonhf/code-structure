package ports

type API interface {
	GetSelection() (int, error)
}
