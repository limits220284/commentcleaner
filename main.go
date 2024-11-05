package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	StatusString = iota
	StatusBlock
	StatusPre
	StatusBlockEnd
)

const (
	SymbolSlash    = '/'
	SymbolAsterisk = '*'
	SymbolSharp    = '#'
)

func removeCommentsForSlash(source []string) []string {
	var cur string
	status := StatusString
	var res []string

	for _, s := range source {
		if s == "" {
			res = append(res, s)
		}
		for i := 0; i < len(s); i++ {
			ch := s[i]

			if status == StatusString {
				if ch == SymbolSlash {
					status = StatusPre
				} else {
					cur += string(ch)
				}
			} else if status == StatusPre {
				if ch == SymbolSlash {
					status = StatusString
					break
				} else if ch == SymbolAsterisk {
					status = StatusBlockEnd
				} else {
					status = StatusString
					cur += string(SymbolSlash) + string(ch)
				}
			} else if status == StatusBlock {
				if ch == SymbolAsterisk {
					status = StatusBlockEnd
				}
			} else if status == StatusBlockEnd {
				if ch == SymbolSlash {
					status = StatusString
				} else if ch != SymbolAsterisk {
					status = StatusBlock
				}
			}
		}

		if status == StatusPre {
			cur += string(SymbolSlash)
			status = StatusString
		} else if status == StatusBlockEnd {
			status = StatusBlock
		}

		if len(cur) != 0 && status == StatusString {
			if strings.TrimSpace(cur) != "" {
				res = append(res, cur)
			}
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
