package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	rootPath := os.Getenv("ROOT_PATH")
	fmt.Println(rootPath)

	archiveId := "baba-board"

	entries, err := os.ReadDir(filepath.Join(rootPath, archiveId))
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}

}
