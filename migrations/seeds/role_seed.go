package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Revprm/go-fp-pbkk/entity"
	"gorm.io/gorm"
)

func ListRoleSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/roles.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listRole []entity.Role
	if err := json.Unmarshal(jsonData, &listRole); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Role{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Role{}); err != nil {
			return err
		}
	}

	for _, data := range listRole {
		var role entity.Role
		err := db.Where(&entity.Role{ID: data.ID}).First(&role).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&role, "id = ?", data.ID).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
