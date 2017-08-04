package models

import (
	"log"

	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

var migrations = []*gormigrate.Migration{
	// you migrations here
	{
		ID: "create_score",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Scores{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable(&Scores{}).Error
		},
	},
	{
		ID: "create_score_record",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&ScoresRecord{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable(&ScoresRecord{}).Error
		},
	},
}

// Migrate do migrate for service
func Migrate(db *gorm.DB) {

	options := gormigrate.DefaultOptions
	options.IDColumnSize = 128
	m := gormigrate.New(db, options, migrations)

	err := m.Migrate()
	if err == nil {
		log.Printf("Migration did run successfully")
	} else {
		log.Printf("Could not migrate: %v", err)
	}

}