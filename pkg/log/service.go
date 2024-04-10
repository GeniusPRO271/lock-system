package log

import (
	"gorm.io/gorm"
)

type LogService interface {
}

type LogServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}
