package src

import (
	"reflect"
	"testing"
)

// Test RemoveCommentsForSlash function
func TestRemoveCommentsForSlash(t *testing.T) {
	source := []string{
		"// This is a single-line comment",
		"int a = 0; // Initialize variable",
		"/* This is a",
		"multi-line comment */",
		"printf(\"Hello, World!\");",
		"/* Comment start",
		"Still in comment",
		"Comment end */",
		"return a;",
	}

	expected := []string{
		"int a = 0; ",
		"printf(\"Hello, World!\");",
		"return a;",
	}

	result := RemoveCommentsForSlash(source)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveCommentsForSlash test failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

// Test RemoveCommentsForHash function
func TestRemoveCommentsForHash(t *testing.T) {
	// Python language test
	source := []string{
		"# -*- coding: utf-8 -*-",
		"'''",
		"This is a multi-line comment",
		"Still in comment",
		"'''",
		"def greet():",
		"    print(\"Hello, World!\")  # Print greeting",
		"",
		"greet()",
	}
	expected := []string{
		"def greet():",
		"    print(\"Hello, World!\")  ",
		"",
		"greet()",
	}

	result := RemoveCommentsForHash(source)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveCommentsForHash test failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

// Test RemoveCommentsForDash function (for SQL, Lua)
func TestRemoveCommentsForDash(t *testing.T) {
	source := []string{
		"-- This is a single-line comment",
		"SELECT * FROM users; -- Get all users",
		"/* Start multi-line comment",
		"Still in comment",
		"End multi-line comment */",
		"INSERT INTO users VALUES ('Alice', 30);",
	}

	expected := []string{
		"SELECT * FROM users; ",
		"INSERT INTO users VALUES ('Alice', 30);",
	}

	result := RemoveCommentsForDash(source)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveCommentsForDash test failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

// Test RemoveCommentsForPercent function (for MATLAB/Octave)
func TestRemoveCommentsForPercent(t *testing.T) {
	source := []string{
		"% This is a single-line comment",
		"x = 1; % Initialize x",
		"%{",
		"This is a multi-line comment",
		"Still in comment",
		"%}",
		"y = x + 1;",
	}

	expected := []string{
		"x = 1; ",
		"y = x + 1;",
	}

	result := RemoveCommentsForPercent(source)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveCommentsForPercent test failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

// Test RemoveCommentsForSemicolon function (for Assembly)
func TestRemoveCommentsForSemicolon(t *testing.T) {
	source := []string{
		"; This is a single-line comment",
		"MOV AX, BX ; Move BX to AX",
		"ADD AX, 1 ; Increment AX",
	}

	expected := []string{
		"MOV AX, BX ",
		"ADD AX, 1 ",
	}

	result := RemoveCommentsForSemicolon(source)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveCommentsForSemicolon test failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

// Test RemoveCommentsForHTML function (for HTML/XML)
func TestRemoveCommentsForHTML(t *testing.T) {
	source := []string{
		"<!-- This is a comment -->",
		"<html>",
		"<head><!-- Head comment --></head>",
		"<body>",
		"<!-- Start multi-line comment",
		"Still in comment",
		"End multi-line comment -->",
		"<p>Hello, World!</p>",
		"</body>",
		"</html>",
	}

	expected := []string{
		"<html>",
		"<head></head>",
		"<body>",
		"<p>Hello, World!</p>",
		"</body>",
		"</html>",
	}

	result := RemoveCommentsForHTML(source)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveCommentsForHTML test failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

// Test FileType function
func TestFileType(t *testing.T) {
	tests := []struct {
		fileName     string
		expectedType Language
		expectedOk   bool
	}{
		{"main.go", Go{}, true},
		{"program.java", Java{}, true},
		{"module.py", Python{}, true},
		{"script.sh", Shell{}, true},
		{"code.pl", Perl{}, true},
		{"app.rb", Ruby{}, true},
		{"analysis.r", RLang{}, true},
		{"script.ps1", PowerShell{}, true},
		{"compute.jl", Julia{}, true},
		{"matrix.m", Matlab{}, true},
		{"query.sql", SQL{}, true},
		{"game.lua", Lua{}, true},
		{"assembly.asm", Assembly{}, true},
		{"index.html", HTML{}, true},
		{"app.js", JavaScript{}, true},
		{"program.cs", CSharp{}, true},
		{"unknown.xyz", nil, false},
	}

	for _, test := range tests {
		lang, ok := FileType(test.fileName)
		if ok != test.expectedOk {
			t.Errorf("FileType(%s) expected ok: %v, got: %v", test.fileName, test.expectedOk, ok)
		}
		if ok && reflect.TypeOf(lang) != reflect.TypeOf(test.expectedType) {
			t.Errorf("FileType(%s) returned incorrect type. Expected %T, got %T", test.fileName, test.expectedType, lang)
		}
	}
}

// Comprehensive test using different language processors
func TestLanguageProcessors(t *testing.T) {
	// Define source code and expected results for different languages
	type testCase struct {
		source   []string
		expected []string
		lang     Language
	}

	testCases := []testCase{
		{
			// Go language test
			source: []string{
				"// Package main",
				"package main",
				"/* Import fmt package */",
				"import \"fmt\"",
				"func main() {",
				"    fmt.Println(\"Hello, World!\") // Print message",
				"}",
			},
			expected: []string{
				"package main",
				"import \"fmt\"",
				"func main() {",
				"    fmt.Println(\"Hello, World!\") ",
				"}",
			},
			lang: Go{},
		},
		{
			// Python language test
			source: []string{
				"# -*- coding: utf-8 -*-",
				"'''",
				"This is a multi-line comment",
				"Still in comment",
				"'''",
				"def greet():",
				"    print(\"Hello, World!\")  # Print greeting",
				"",
				"greet()",
			},
			expected: []string{
				"def greet():",
				"    print(\"Hello, World!\")  ",
				"",
				"greet()",
			},
			lang: Python{},
		},
		{
			// SQL language test
			source: []string{
				"-- Create table",
				"CREATE TABLE users (",
				"    id INT PRIMARY KEY, /* User ID */",
				"    name VARCHAR(100) -- Username",
				");",
			},
			expected: []string{
				"CREATE TABLE users (",
				"    id INT PRIMARY KEY, ",
				"    name VARCHAR(100) ",
				");",
			},
			lang: SQL{},
		},
		{
			// HTML test
			source: []string{
				"<!-- HTML document start -->",
				"<!DOCTYPE html>",
				"<html>",
				"<head>",
				"    <title>Sample Page</title>",
				"</head>",
				"<body>",
				"    <!-- Main content -->",
				"    <h1>Hello, World!</h1>",
				"</body>",
				"</html>",
			},
			expected: []string{
				"<!DOCTYPE html>",
				"<html>",
				"<head>",
				"    <title>Sample Page</title>",
				"</head>",
				"<body>",
				"    <h1>Hello, World!</h1>",
				"</body>",
				"</html>",
			},
			lang: HTML{},
		},
	}

	for _, test := range testCases {
		result := test.lang.RemoveComments(test.source)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Language processor test failed for %T.\nExpected: %v\nGot: %v", test.lang, test.expected, result)
		}
	}
}
