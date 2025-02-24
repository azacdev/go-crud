package initializers

import "github.com/azacdev/go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
}
