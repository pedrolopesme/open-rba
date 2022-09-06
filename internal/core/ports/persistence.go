package ports

type Repository interface {
	Insert(key string, row interface{}) error
}
