package ports

type Repository interface {
	Insert(row interface{}) error
}
