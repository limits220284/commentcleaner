package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/limits220284/commentcleaner/src"
)

func readFile(filePath string) (ans []string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ans = append(ans, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return
}

func writeToFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func processFile(filePath string) {
	content := readFile(filePath)
	result := src.RemoveCommentsForSlash(content)
	writeToFile("todo.txt", result)
}

func isTargetFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".go")
}

func processFiles(filesPath string) error {
	err := filepath.Walk(filesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if isTargetFile(info.Name()) {
			log.Printf("Processing file: %s\n", path)
			processedContent := src.RemoveCommentsForSlash(readFile(path))
			err := writeToFile(path, processedContent)
			if err != nil {
				fmt.Printf("Error processing file %s: %s\n", path, err)
			}
		}

		return nil
	})

	return err
}

func main() {
	// 一个比较好的做法是，直接拷贝一份，然后创建，然后对拷贝的进行递归处理即可
	filePath := "todo.go"
	processFile(filePath)
	filesPath := "./test_path"
	processFiles(filesPath)
}
