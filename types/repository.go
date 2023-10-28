package types

type Repository interface {
	Save(data any) (any, error)
	Update(id string, data any) (any, error)
	Delete(id string) (any, error)
	FindById(id string) (any, error)
	FindOne(filter map[string]any) (any, error)
	FindAll(filter map[string]any) (any, error)
}
