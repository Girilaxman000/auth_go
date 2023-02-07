package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// capital letter for using in other files
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env files")
	}
}
