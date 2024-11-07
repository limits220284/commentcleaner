package src

import "strings"

type Language interface {
	RemoveComments(source []string) []string
}

func FileType(fileName string) (Language, bool) {
	typ := strings.Split(fileName, ".")[1]
	switch typ {
	case "go":
		return Go{}, true
	case "java":
		return Java{}, true
	case "cpp", "h", "c":
		return Cpp{}, true
	case "py":
		return Python{}, true
	}
	return nil, false
}

type Go struct {
}

func (Go) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}

type Java struct {
}

func (Java) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}

type Python struct {
}

func (Python) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

type Cpp struct {
}

func (Cpp) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}
