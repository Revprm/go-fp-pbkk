package migrations

import (
	"github.com/Revprm/go-fp-pbkk/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
	); err != nil {
		return err
	}

	return nil
}
