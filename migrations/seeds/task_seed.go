package seeds

import (
	"encoding/json"
	"errors"
	"fmt" // Add for debugging
	"io"
	"os"

	"github.com/Revprm/go-fp-pbkk/entity"
	"gorm.io/gorm"
)

func ListTaskSeeder(db *gorm.DB) error {
	fmt.Println("Starting the seeding process...") // Debug start

	// Open the JSON file
	jsonFile, err := os.Open("./migrations/json/tasks.json")
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err) // Debug error
		return err
	}
	defer jsonFile.Close()

	// Read JSON data
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err) // Debug error
		return err
	}

	// Parse JSON data
	var listTask []entity.Task
	if err := json.Unmarshal(jsonData, &listTask); err != nil {
		fmt.Printf("Error unmarshaling JSON data: %v\n", err) // Debug error
		return err
	}
	fmt.Printf("Loaded %d tasks from JSON file.\n", len(listTask)) // Debug task count

	// Check if the table exists
	hasTable := db.Migrator().HasTable(&entity.Task{})
	if !hasTable {
		fmt.Println("Task table does not exist. Creating table...") // Debug table creation
		if err := db.Migrator().CreateTable(&entity.Task{}); err != nil {
			fmt.Printf("Error creating Task table: %v\n", err) // Debug error
			return err
		}
	} else {
		fmt.Println("Task table already exists.") // Debug confirmation
	}

	// Insert tasks into the database
	for i, data := range listTask {
		fmt.Printf("Processing task %d: %+v\n", i+1, data) // Debug task details

		var t entity.Task
		err := db.Where("id = ?", data.ID).First(&t).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("Task with ID %v not found. Creating new record...\n", data.ID) // Debug not found
				if err := db.Create(&data).Error; err != nil {
					fmt.Printf("Error creating task with ID %v: %v\n", data.ID, err) // Debug error
					return err
				}
				fmt.Printf("Task with ID %v created successfully.\n", data.ID) // Debug success
			} else {
				fmt.Printf("Error querying task with ID %v: %v\n", data.ID, err) // Debug error
				return err
			}
		} else {
			fmt.Printf("Task with ID %v already exists. Skipping...\n", data.ID) // Debug skip
		}
	}

	fmt.Println("Seeding process completed successfully.") // Debug end
	return nil
}
