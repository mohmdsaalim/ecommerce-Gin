package repositories

type Repository interface {
	Insert(data interface{}) error
	FindByID(model interface{}, id interface{}) error
	UpdateByID(model interface{}, id interface{}, data interface{}) error
	UpdateFields(model interface{}, id interface{}, fields map[string]interface{}) error
	Delete(model interface{}, query string, args ...interface{}) error
	FindOne(dest interface{}, query string, preloads []string, args ...interface{}) error
	FindAll(dest interface{}, query string, order string, preloads []string, args ...interface{}) error


	Count(model interface{}, query string, args ...interface{}) (int64, error)
	Sum(model interface{}, column string, query string, args ...interface{}) (float64, error)
	FindWithLimit(dest interface{}, query string, order string, limit int, preloads []string, args ...interface{}) error
}
