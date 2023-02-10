package infrastructure

import (
	"CRUD/src/interfaces/database"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

// Create implements database.SqlHandler
func (handler *SqlHandler) Create(object interface{}) {
	handler.db.Create(object)
}

// DeleteById implements database.SqlHandler
func (handler *SqlHandler) DeleteById(object interface{}, id string) {

	handler.db.Delete(object, id)
}

// FindAll implements database.SqlHandler
func (handler *SqlHandler) FindAll(object interface{}) {
	handler.db.Find(object)
}

func NewSqlHandler(sc StorageConfig) database.SqlHandler {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}
