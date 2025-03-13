package api

import (
	"log"
	"os"
	"path"
)

func GetSchema() string {
	projectPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	schemaPath := path.Join(projectPath, "cmd", "kafka", "http", "schema.avsc")

	file, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
