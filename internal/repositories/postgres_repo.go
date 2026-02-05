package repositories

import (
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
	"gorm.io/gorm"
)

type PgSQLRepository struct {
	db *gorm.DB
}
// gorm connection 
func NewPgSQLRepository() Repository {
	return &PgSQLRepository{
		db: database.DB,
	}

}

func (r *PgSQLRepository) Insert(data interface{}) error {
	return r.db.Debug().Create(data).Error
}

func (r *PgSQLRepository) FindByID(model interface{}, id interface{}) error {
	return r.db.First(model, id).Error
}

func (r *PgSQLRepository) UpdateByID(model interface{}, id interface{}, data interface{}) error {
	return r.db.Model(model).Where("id = ?", id).Updates(data).Error
}

func (r *PgSQLRepository) UpdateFields(model interface{}, id interface{}, fields map[string]interface{}) error {
	return r.db.Model(model).Where("id = ?", id).Updates(fields).Error
}

func (r *PgSQLRepository) Delete(model interface{}, id interface{}) error {
	return r.db.Delete(model, id).Error
}

func (r *PgSQLRepository) FindOne(dest interface{}, query string, args ...interface{}) error {
	return r.db.Where(query, args...).First(dest).Error
}

func (r *PgSQLRepository) FindAll(dest interface{}, query string, args ...interface{}) error {
	return r.db.Where(query, args...).Find(dest).Error
}

func (r *PgSQLRepository) FindProductWithCategory(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.
		Where("is_available = ?", true).
		First(&product, id).Error
	return &product, err
}

func (r *PgSQLRepository) FindAllProductsWithCategory() ([]models.Product, error) {
	var products []models.Product
	err := r.db.
		Where("is_available = ?", true).
		Order("created_at DESC").
		Find(&products).Error
	return products, err
}

func (r *PgSQLRepository) FindProductsByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.
		Where("category = ? AND is_available = ?", category, true).
		Order("created_at DESC").
		Find(&products).Error
	return products, err
}

func (r *PgSQLRepository) FindProductsBySubCategory(category, subCategory string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.
		Where("category = ? AND sub_category = ? AND is_available = ?", category, subCategory, true).
		Order("created_at DESC").
		Find(&products).Error
	return products, err
}


// // Get all products
// products, err := repo.FindAllProductsWithCategory()

// // Get all kits
// kits, err := repo.FindProductsByCategory("kits")

// // Get home kits only
// homeKits, err := repo.FindProductsBySubCategory("kits", "home")

// // Get lifestyle hoodies
// hoodies, err := repo.FindProductsBySubCategory("lifestyle", "hoodie")