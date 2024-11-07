package src

import "strings"

// Language interface definition
type Language interface {
	RemoveComments(source []string) []string
}

// Return the corresponding language processor based on the file extension
func FileType(fileName string) (Language, bool) {
	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		return nil, false
	}
	ext := parts[len(parts)-1]
	switch ext {
	case "go":
		return Go{}, true
	case "java":
		return Java{}, true
	case "cpp", "h", "c":
		return Cpp{}, true
	case "py":
		return Python{}, true
	case "sh", "bash", "zsh":
		return Shell{}, true
	case "pl":
		return Perl{}, true
	case "rb":
		return Ruby{}, true
	case "r":
		return RLang{}, true
	case "ps1":
		return PowerShell{}, true
	case "jl":
		return Julia{}, true
	case "m":
		return Matlab{}, true
	case "sql":
		return SQL{}, true
	case "lua":
		return Lua{}, true
	case "asm":
		return Assembly{}, true
	case "html", "htm", "xml":
		return HTML{}, true
	case "js":
		return JavaScript{}, true
	case "cs":
		return CSharp{}, true
	}
	return nil, false
}

// Define the structure for each language and implement the RemoveComments method

// Go language
type Go struct{}

func (Go) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}

// Java language
type Java struct{}

func (Java) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}

// C/C++ language
type Cpp struct{}

func (Cpp) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}

// Python language
type Python struct{}

func (Python) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

// Shell script (uses # for comments)
type Shell struct{}

func (Shell) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

// Perl language
type Perl struct{}

func (Perl) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

// Ruby language
type Ruby struct{}

func (Ruby) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

// R language
type RLang struct{}

func (RLang) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

// PowerShell script
type PowerShell struct{}

func (PowerShell) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

// Julia language
type Julia struct{}

func (Julia) RemoveComments(source []string) []string {
	return RemoveCommentsForHash(source)
}

// MATLAB/Octave language (uses % for comments)
type Matlab struct{}

func (Matlab) RemoveComments(source []string) []string {
	return RemoveCommentsForPercent(source)
}

// SQL language (uses -- and /* */ for comments)
type SQL struct{}

func (SQL) RemoveComments(source []string) []string {
	return RemoveCommentsForDash(source)
}

// Lua language (uses -- and --[[ ]] for comments)
type Lua struct{}

func (Lua) RemoveComments(source []string) []string {
	return RemoveCommentsForDash(source)
}

// Assembly language (uses ; for comments)
type Assembly struct{}

func (Assembly) RemoveComments(source []string) []string {
	return RemoveCommentsForSemicolon(source)
}

// HTML/XML (uses <!-- --> for comments)
type HTML struct{}

func (HTML) RemoveComments(source []string) []string {
	return RemoveCommentsForHTML(source)
}

// JavaScript language
type JavaScript struct{}

func (JavaScript) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}

// C# language
type CSharp struct{}

func (CSharp) RemoveComments(source []string) []string {
	return RemoveCommentsForSlash(source)
}
