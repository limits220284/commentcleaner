package main

import (
	"bufio"
	"fmt"
	"os"
)

func removeCommentsForSlash(source []string) []string {
	var cur string
	status := "str"
	var res []string

	for _, s := range source {
		for i := 0; i < len(s); i++ {
			ch := s[i]

			if status == "str" {
				if ch == '/' {
					status = "pre"
				} else {
					cur += string(ch)
				}
			} else if status == "pre" {
				if ch == '/' {
					status = "str"
					break
				} else if ch == '*' {
					status = "block"
				} else {
					status = "str"
					cur += "/" + string(ch)
				}
			} else if status == "block" {
				if ch == '*' {
					status = "block_end_pre"
				}
			} else if status == "block_end_pre" {
				if ch == '/' {
					status = "str"
				} else if ch != '*' {
					status = "block"
				}
			}
		}

		if status == "pre" {
			cur += "/"
			status = "str"
		} else if status == "block_end_pre" {
			status = "block"
		}

		if len(cur) != 0 && status == "str" {
			res = append(res, cur)
			cur = ""
		}
	}

	return res
}

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

func main() {
	filePath := "todo.go"
	content := readFile(filePath)
	result := removeCommentsForSlash(content)
	writeToFile("todo.txt", result)
}
