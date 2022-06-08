package services

type DoService interface {
	Do() (bool, error)
}
