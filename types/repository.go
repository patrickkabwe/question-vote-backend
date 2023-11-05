package types

type Repository interface {
	Create(data any) (any, error)
	Update(id int64, data any) (any, error)
	Delete(id string) (any, error)
	FindById(id string) (any, error)
	FindOne(filter map[string]any) (any, error)
	FindAll(filter map[string]any) (any, error)
}
