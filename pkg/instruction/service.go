package instruction

import (
	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"gorm.io/gorm"
)

type InstructionService interface {
}

type InstructionServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

func (s *InstructionServiceImpl) PostLog(log model.Log) error {

	if result := s.Db.Create(&log); result.Error != nil {
		return result.Error
	}

	return nil
}
