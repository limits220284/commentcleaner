package src

import "strings"

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
