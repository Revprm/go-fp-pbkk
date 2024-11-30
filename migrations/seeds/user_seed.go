package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Revprm/go-fp-pbkk/entity"
	"gorm.io/gorm"
)

func ListUserSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/users.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, _ := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listUser []entity.User
	if err := json.Unmarshal(jsonData, &listUser); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	for _, data := range listUser {
		var user entity.User
		err := db.Where(&user, "email = ?", data.Email).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
