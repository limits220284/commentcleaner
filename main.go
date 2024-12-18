package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/limits220284/commentcleaner/src"
	"github.com/limits220284/commentcleaner/utils"
)

func processFiles(filesPath string) error {
	abFilesPath, _ := filepath.Abs(filesPath)
	processedFilesPath := filepath.Dir(abFilesPath) + "\\processed_" + filepath.Base(filesPath)
	log.Println(processedFilesPath)
	err := os.CopyFS(processedFilesPath, os.DirFS(filesPath))
	if err != nil {
		log.Fatal("copy files from source path failed!")
	}
	err = filepath.Walk(processedFilesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if worker, ok := src.FileType(info.Name()); ok {
			log.Printf("Processing file: %s\n", path)
			processedContent := worker.RemoveComments(utils.ReadFile(path))
			err := utils.WriteToFile(path, processedContent)
			if err != nil {
				fmt.Printf("Error processing file %s: %s\n", path, err)
			}
		}

		return nil
	})

	return err
}

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: go run main.go <filePath>")
		return
	}
	filesPath := os.Args[1]
	processFiles(filesPath)
}
