package src

// Status constants
const (
	StatusString      = iota // Normal code state
	StatusBlock              // Multi-line comment state
	StatusPre                // Pre-processing state (used to handle comment start symbol)
	StatusBlockEnd           // Multi-line comment end state
	StatusLineComment        // Single-line comment state
)

// Symbol constants
const (
	SymbolSlash            = '/'
	SymbolAsterisk         = '*'
	SymbolSharp            = '#'
	SymbolQuote            = '"'
	SymbolSingleQuote      = '\''
	SymbolPercent          = '%'
	SymbolSemicolon        = ';'
	SymbolDash             = '-'
	SymbolLessThan         = '<'
	SymbolGreaterThan      = '>'
	SymbolLeftBrace        = '{'
	SymbolRightBrace       = '}'
	SymbolHTMLCommentStart = "<!--"
	SymbolHTMLCommentEnd   = "-->"
)
