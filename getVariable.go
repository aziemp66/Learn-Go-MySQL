package gomysql

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	//check current directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// load .env file
	godotenv.Load(dir + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
