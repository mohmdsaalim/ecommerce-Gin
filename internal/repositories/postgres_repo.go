package repositories

import (
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
	return r.db.Debug().Create(data).Error // creating data add data into db ✅
}

func (r *PgSQLRepository) FindByID(model interface{}, id interface{}) error {
	return r.db.First(model, id).Error
} // take one specific data user1 or user2  ✅

//	func (r *PgSQLRepository) UpdateByID(model interface{}, id interface{}, data interface{}) error {
//		return r.db.Model(model).Where("id = ?", id).Updates(data).Error
//	}
func (r *PgSQLRepository) UpdateByID(model interface{}, id interface{}, data interface{}) error {
	return r.db.
		Model(model).
		Where("id = ?", id).
		Omit("Variants").
		Updates(data).Error
}

func (r *PgSQLRepository) UpdateFields(model interface{}, id interface{}, fields map[string]interface{}) error {
	return r.db.Model(model).Where("id = ?", id).Updates(fields).Error
}

func (r *PgSQLRepository) Delete(model interface{}, query string, args ...interface{}) error {
	return r.db.Where(query, args...).Delete(model).Error
}

func (r *PgSQLRepository) FindOne(dest interface{}, query string, preloads []string, args ...interface{}) error {
	db := r.db
	if query != "" {
		db = db.Where(query, args...)
	}
	for _, p := range preloads {
		db = db.Preload(p)
	}
	return db.First(dest).Error
} //

func (r *PgSQLRepository) FindAll(dest interface{}, query string, order string, preloads []string, args ...interface{}) error {
	db := r.db
	if query != "" {
		db = db.Where(query, args...)
	}
	if order != "" {
		db = db.Order(order)
	}
	for _, p := range preloads {
		db = db.Preload(p)
	}
	return db.Find(dest).Error
}

// FindWithPagination retrieves data with both limit and offset (for pages)
func (r *PgSQLRepository) FindWithPagination(
	dest interface{},
	query string,
	order string,
	limit int,
	offset int,
	preloads []string,
	args ...interface{},
) error {
	db := r.db

	// apply where clause if query is provided
	if query != "" {
		db = db.Where(query, args...)
	}

	// apply ordering if provided
	if order != "" {
		db = db.Order(order)
	}

	// limit is how many items per page
	if limit > 0 {
		db = db.Limit(limit)
	}

	// offset is how many items to skip (page-1 * limit)
	if offset >= 0 {
		db = db.Offset(offset)
	}

	// preload any relations (like Variants for Products)
	for _, p := range preloads {
		db = db.Preload(p)
	}

	// execute the query
	return db.Find(dest).Error
}

func (r *PgSQLRepository) Count(model interface{}, query string, args ...interface{}) (int64, error) {
	var count int64
	db := r.db.Model(model)
	if query != "" {
		db = db.Where(query, args...)
	}
	err := db.Count(&count).Error
	return count, err
}

func (r *PgSQLRepository) Sum(model interface{}, column string, query string, args ...interface{}) (float64, error) {
	var total float64

	db := r.db.Model(model)

	if query != "" {
		db = db.Where(query, args...)
	}

	err := db.Select("COALESCE(SUM(" + column + "),0)").Scan(&total).Error
	return total, err
}

func (r *PgSQLRepository) FindWithLimit(
	dest interface{},
	query string,
	order string,
	limit int,
	preloads []string,
	args ...interface{},
) error {

	db := r.db

	if query != "" {
		db = db.Where(query, args...)
	}

	if order != "" {
		db = db.Order(order)
	}

	if limit > 0 {
		db = db.Limit(limit)
	}

	for _, p := range preloads {
		db = db.Preload(p)
	}

	return db.Find(dest).Error
}
