package repositories

type Repository interface {
	Insert(data interface{}) error
	FindByID(model interface{}, id interface{}) error
	UpdateByID(model interface{}, id interface{}, data interface{}) error
	UpdateFields(model interface{}, id interface{}, fields map[string]interface{}) error
	Delete(model interface{}, id interface{}) error
}
