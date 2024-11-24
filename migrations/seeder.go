package migrations

import (
	"github.com/Revprm/go-fp-pbkk/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListRoleSeeder(db); err != nil {
		return err
	}

	return nil
}
