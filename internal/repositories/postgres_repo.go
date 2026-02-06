package repositories

import (
	// "github.com/mohmdsaalim/ecommerce-Gin/internal/models"
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

func (r *PgSQLRepository) FindAll(dest interface{}, query string, order string, preloads []string, args ...interface{}) error {
	db := r.db
	if query != "" { db = db.Where(query, args...) }
	if order != "" { db = db.Order(order) }
	for _, preload := range preloads { db = db.Preload(preload) }
	return db.Find(dest).Error
}
