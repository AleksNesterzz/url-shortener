package storage

type Storage interface {
	Get(url string) (string, error)
	Create(url string) (string, error)
	Delete(url string) error
}
