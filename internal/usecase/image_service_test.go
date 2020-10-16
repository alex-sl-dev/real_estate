package usecase

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}

}
func TestImageService_Resize(t *testing.T) {

	t.Error("pass")
}
