package repositories

// import "github.com/mohmdsaalim/ecommerce-Gin/internal/models"

type Repository interface {
	Insert(data interface{}) error
	FindByID(model interface{}, id interface{}) error
	UpdateByID(model interface{}, id interface{}, data interface{}) error
	UpdateFields(model interface{}, id interface{}, fields map[string]interface{}) error
	Delete(model interface{}, id interface{}) error
	FindOne(dest interface{}, query string, args ...interface{}) error
	// FindAll(dest interface{}, query string, args... interface{}) error
	FindAll(dest interface{}, query string, order string, preloads []string, args ...interface{}) error
	
	// Products specific methods
	// GetAllProducts()([]models.Product, error)
}
