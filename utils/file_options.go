package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filePath string) (ans []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ans = append(ans, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}
	return
}

func WriteToFile(filePath string, lines []string) error {
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
