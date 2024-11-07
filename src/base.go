package src

import "strings"

// remove // and /* */ (like go、java、cpp)
func RemoveCommentsForSlash(source []string) []string {
	var cur string
	status := StatusString
	var res []string

	for _, s := range source {
		if s == "" {
			res = append(res, s)
			continue
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

// remove # and ''' ''' or """ """（like Python）
func RemoveCommentsForHash(source []string) []string {
	var result []string
	state := StatusString

	for _, line := range source {
		if line == "" {
			result = append(result, line)
			continue
		}
		i := 0
		var buffer strings.Builder
		for i < len(line) {
			switch state {
			case StatusString:
				if line[i] == SymbolSharp {
					state = StatusLineComment
				} else if i <= len(line)-3 && ((line[i] == SymbolQuote && line[i+1] == SymbolQuote && line[i+2] == SymbolQuote) ||
					(line[i] == SymbolSingleQuote && line[i+1] == SymbolSingleQuote && line[i+2] == SymbolSingleQuote)) {
					state = StatusBlock
					i += 2
				} else {
					buffer.WriteByte(line[i])
				}
			case StatusLineComment:
				i = len(line) - 1
			case StatusBlock:
				if i <= len(line)-3 && ((line[i] == SymbolQuote && line[i+1] == SymbolQuote && line[i+2] == SymbolQuote) ||
					(line[i] == SymbolSingleQuote && line[i+1] == SymbolSingleQuote && line[i+2] == SymbolSingleQuote)) {
					state = StatusString
					i += 2
				}
			}
			i++
		}

		if (state == StatusString || state == StatusLineComment) && strings.TrimSpace(buffer.String()) != "" {
			result = append(result, buffer.String())
		}

		if state == StatusLineComment {
			state = StatusString
		}
	}

	return result
}

// remove -- and /* */ （like SQL、Lua）
func RemoveCommentsForDash(source []string) []string {
	var result []string
	state := StatusString

	for _, line := range source {
		i := 0
		var buffer strings.Builder
		for i < len(line) {
			switch state {
			case StatusString:
				if i < len(line)-1 && line[i] == SymbolDash && line[i+1] == SymbolDash {
					state = StatusLineComment
					i = len(line)
				} else if i < len(line)-1 && line[i] == SymbolSlash && line[i+1] == SymbolAsterisk {
					state = StatusBlock
					i++
				} else {
					buffer.WriteByte(line[i])
				}
			case StatusBlock:
				if i < len(line)-1 && line[i] == SymbolAsterisk && line[i+1] == SymbolSlash {
					state = StatusString
					i++
				}
			}
			i++
		}

		if (state == StatusString || state == StatusLineComment) && strings.TrimSpace(buffer.String()) != "" {
			result = append(result, buffer.String())
		}

		if state == StatusLineComment {
			state = StatusString
		}
	}

	return result
}

// RemoveCommentsForPercent removes comments starting with '%' and multi-line comments enclosed between '%{' and '%}'
func RemoveCommentsForPercent(source []string) []string {
	var result []string
	inBlockComment := false

	for _, line := range source {
		var buffer strings.Builder
		i := 0
		for i < len(line) {
			if !inBlockComment {
				if i+1 < len(line) && line[i] == SymbolPercent && line[i+1] == SymbolLeftBrace {
					inBlockComment = true
					i += 1 // Skip '{'
				} else if line[i] == SymbolPercent {
					// Single-line comment, skip the rest of the line
					break
				} else {
					buffer.WriteByte(line[i])
				}
			} else {
				// Inside multi-line comment
				if i+1 < len(line) && line[i] == SymbolPercent && line[i+1] == SymbolRightBrace {
					inBlockComment = false
					i += 1 // Skip '}'
				}
				// Skip characters inside multi-line comment
			}
			i++
		}

		// Add the line to the result if we're not inside a block comment and the line is not empty
		if !inBlockComment && strings.TrimSpace(buffer.String()) != "" {
			result = append(result, buffer.String())
		}
	}

	return result
}

// remove ;（like Assembly）
func RemoveCommentsForSemicolon(source []string) []string {
	var result []string

	for _, line := range source {
		var buffer strings.Builder
		i := 0
		for i < len(line) {
			if line[i] == SymbolSemicolon {
				break
			} else {
				buffer.WriteByte(line[i])
			}
			i++
		}

		if strings.TrimSpace(buffer.String()) != "" {
			result = append(result, buffer.String())
		}
	}

	return result
}

// remove  <!-- -->（like HTML、XML）
func RemoveCommentsForHTML(source []string) []string {
	var result []string
	state := StatusString

	for _, line := range source {
		i := 0
		var buffer strings.Builder
		for i < len(line) {
			switch state {
			case StatusString:
				if i <= len(line)-4 && line[i:i+4] == SymbolHTMLCommentStart {
					state = StatusBlock
					i += 3
				} else {
					buffer.WriteByte(line[i])
				}
			case StatusBlock:
				if i <= len(line)-3 && line[i:i+3] == SymbolHTMLCommentEnd {
					state = StatusString
					i += 2
				}
			}
			i++
		}

		if state == StatusString && strings.TrimSpace(buffer.String()) != "" {
			result = append(result, buffer.String())
		}
	}

	return result
}
