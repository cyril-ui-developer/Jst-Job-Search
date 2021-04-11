package db

import (
	"github.com/jinzhu/gorm"
	"github/cyril-ui-developer/JstJobSearch/internal/jobs"
)

//MigrateDB  - migrates db and create jobs table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&jobs.Job{}); result.Error != nil {
		return result.Error
	}

	return nil
}