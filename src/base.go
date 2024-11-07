package src

import "strings"

func RemoveCommentsForSlash(source []string) []string {
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

func RemoveCommentsForHash(source []string) []string {
	var result []string
	state := StatusString

	for _, line := range source {
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
