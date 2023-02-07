package migrate

import (
	"github.com/Girilaxman000/auth_go/database"
	"github.com/Girilaxman000/auth_go/models"
)

func SyncDatabase() {
	database.DB.AutoMigrate(&models.User{})
}
