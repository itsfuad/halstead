package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"flag"
)

type LanguageRule struct {
	//all regexes for the language
	Name      string
	Operators []string
	Operands  []string
	Skip      []string
}

type Position struct {
	// Line number
	Line int
	// Column number
	Column int
}

type Pattern struct {
	// Regex pattern
	regex *regexp.Regexp
	// Handler function
	handler func(*Lexer, *regexp.Regexp)
}

type Lexer struct {
	// File path
	FilePath string
	// Source code
	SourceCode []byte
	// Current position in the source code
	Position Position
	// Tokens
	Tokens []Token
	// Patterns
	Patterns []Pattern
}

// Fix Lexer implementation:
func (l *Lexer) atEOF() bool {
	return l.Position.Line >= len(string(l.SourceCode))
}

func (l *Lexer) advance(value string) {
	l.Position.Line += len(value)
	if l.Position.Line > len(string(l.SourceCode)) {
		l.Position.Line = len(string(l.SourceCode))
	}
}

func (l *Lexer) at() byte {
	return l.SourceCode[l.Position.Line]
}

func (l *Lexer) remainder() string {
	return string(l.SourceCode[l.Position.Line:])
}

func (l *Lexer) push(token Token) {
	l.Tokens = append(l.Tokens, token)
}

func NewToken(tokenType TokenType, value string, start Position, end Position) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}

var rulesets = map[string]LanguageRule{
	"cpp": {
		Name: "cpp",
		Operators: []string{
			`#include\s*<[^>]+>`, // Include directives
			`\b\w+\s*\([^)]*\)`,  // Function calls
			`[-+*/=<>!&|^%]=?`,   // Basic operators and assignments
			`\+\+|--`,            // Increment/decrement
			`&&|\|\|`,            // Logical operators
			`\b(if|else|while|for|do|switch|case|break|continue|return|new|delete)\b`, // Keywords
			`\b(class|struct|namespace|public|private|protected)\b`,                   // OOP keywords
			`<<|>>`,       // Stream operators
			`::|->|\.|::`, // Scope and member access
			`\[\]`,        // Array subscript
		},
		Operands: []string{
			`"[^"]*"`,            // String literals
			`'[^']*'`,            // Character literals
			`\b\d+(\.\d+)?\b`,    // Numbers (integer and float)
			`\b[a-zA-Z_]\w*\b`,   // Identifiers
			`\btrue\b|\bfalse\b`, // Boolean literals
			`\bnullptr\b`,        // Null pointer
		},
		Skip: []string{
			`\/\/.*`,           // Single line comments
			`\/\*[\s\S]*?\*\/`, // Multi line comments
			`\s+`,              // Whitespace
		},
	},
	"java": {
		Name: "java",
		Operators: []string{
			`[{}()]`, // Braces, parentheses
			`[;,:]`,   // Semicolons, commas
			`\baspect\b|\bpointcut\b|\bexecution\b|\bbefore\b`,
			`\bthisJoinPoint\b`,
			`\b\w+\s*\([^)]*\)`,
			`[-+*/=<>!&|^%]=?`,
			`\+\+|--`,
			`&&|\|\|`,
			`\b(if|else|while|for|do|switch|case|break|continue|return|new)\b`,
			`\b(class|interface|extends|implements|public|private|protected)\b`,
			`\b(try|catch|finally|throw|throws)\b`,
			`->|\.|@`,
			`\[\]`,
		},
		Operands: []string{
			`"[^"]*"`,
			`'[^']*'`,
			`\b\d+(\.\d+)?\b`,
			`\b[a-zA-Z_]\w*\b`,
			`\btrue\b|\bfalse\b`,
			`\bnull\b`,
		},
		Skip: []string{
			`\/\/.*`,
			`\/\*[\s\S]*?\*\/`,
			`\s+`,
		},
	},
	"python": {
		Name: "python",
		Operators: []string{
			`\bimport\b|\bfrom\b`, // Import statements
			`\b\w+\s*\([^)]*\)`,   // Function calls
			`[-+*/=<>!&|^%]=?`,    // Basic operators and assignments
			`\*\*`,                // Power operator
			`and|or|not`,          // Logical operators
			`\b(if|elif|else|while|for|in|break|continue|return|def|class)\b`, // Keywords
			`\b(try|except|finally|raise|with|as)\b`,                          // Exception handling
			`\.|@`,                                                            // Member access, decorators
			`\[\]`,                                                            // List subscript
			`:\s*$`,                                                           // Block delimiter
		},
		Operands: []string{
			`"[^"]*"|'[^']*'`,    // String literals (single/double quotes)
			`\b\d+(\.\d+)?\b`,    // Numbers
			`\b[a-zA-Z_]\w*\b`,   // Identifiers
			`\bTrue\b|\bFalse\b`, // Boolean literals
			`\bNone\b`,           // None literal
		},
		Skip: []string{
			`#.*`,            // Single line comments
			`'''[\s\S]*?'''`, // Multi line comments (triple quotes)
			`"""[\s\S]*?"""`, // Multi line comments (triple double quotes)
			`\s+`,            // Whitespace
		},
	},
}

func createLexer(filename string) *Lexer {

	// Read the source code from the file
	sourceCode, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lang := filename[strings.LastIndex(filename, ".")+1:]

	fmt.Printf("Language: %s\n", lang)

	lex := &Lexer{
		FilePath:   filename,
		SourceCode: sourceCode,
		Position:   Position{Line: 0, Column: 0},
		Tokens:     []Token{},
	}

	rule, ok := rulesets[lang]
	if !ok {
		panic("Language not supported")
	}

	// Create regex patterns
	for _, operator := range rule.Operators {
		pattern := regexp.MustCompile(operator)
		lex.Patterns = append(lex.Patterns, Pattern{regex: pattern, handler: operatorHandler})
	}

	for _, operand := range rule.Operands {
		pattern := regexp.MustCompile(operand)
		lex.Patterns = append(lex.Patterns, Pattern{regex: pattern, handler: operandHandler})
	}

	for _, skip := range rule.Skip {
		pattern := regexp.MustCompile(skip)
		lex.Patterns = append(lex.Patterns, Pattern{regex: pattern, handler: skipHandler})
	}

	return lex
}

func operatorHandler(lex *Lexer, regex *regexp.Regexp) {
	operator := regex.FindString(lex.remainder())
	start := lex.Position
	lex.advance(operator)
	end := lex.Position
	lex.push(NewToken(OPERATOR, operator, start, end))
}

func operandHandler(lex *Lexer, regex *regexp.Regexp) {
	operand := regex.FindString(lex.remainder())
	start := lex.Position
	lex.advance(operand)
	end := lex.Position
	lex.push(NewToken(OPERAND, operand, start, end))
}

// skipHandler processes a token that should be skipped by the lexer.
func skipHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.advance(match)
}

// Tokenize reads the source code from the specified file and tokenizes it.
func Tokenize(filename string) []Token {

	lex := createLexer(filename)
	lex.FilePath = filename

	for !lex.atEOF() {

		matched := false

		for _, pattern := range lex.Patterns {

			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			errStr := fmt.Sprintf("lexer:unexpected charecter: '%c'\n", lex.at())
			panic(errStr)
		}
	}

	return lex.Tokens
}

// Analyze source code with a specific parser
func analyzeSourceCode(filepath string) (Halstead, error) {

	tokens := Tokenize(filepath)

	metrics := GetHalsteadMetrics(tokens)

	return metrics, nil
}

type Halstead struct {
	// Number of distinct operators
	N1 int
	// Number of distinct operands
	N2 int
	// Total number of operators
	n1 int
	// Total number of operands
	n2 int

	// Program vocabulary
	N int
	// Program length
	n int
	// Calculated program length
	Np float64
	// Calculated program volume
	V float64
	// Calculated program difficulty
	D float64
	// Calculated program effort
	E float64
	// Calculated program time
	T float64
	// Calculated program bugs
	B float64
}

func (h *Halstead) Calculate() {
	h.N = h.N1 + h.N2
	h.n = h.n1 + h.n2

	h.Np = float64(h.N1) * (float64(h.N2) / 2)
	h.V = float64(h.N) * (math.Log2(float64(h.N)))
	h.D = (float64(h.N1) / 2) * (float64(h.n2) / float64(h.N2))
	h.E = h.D * h.V
	h.T = h.E / 18   // 18 is the average number of bugs per hour
	h.B = h.V / 3000 // 3000 is the average volume of a bug
}

func (h *Halstead) Print() {
	fmt.Println("--------------------- Halstead Metrics ---------------------")
	fmt.Println("Number of distinct operators \t(N1):", h.N1)
	fmt.Println("Number of distinct operands \t(N2):", h.N2)
	fmt.Println("Total number of operators \t(n1):", h.n1)
	fmt.Println("Total number of operands \t(n2):", h.n2)
	fmt.Println("Program vocabulary \t\t(N):", h.N)
	fmt.Println("Program length \t\t\t(n):", h.n)
	fmt.Println("Calculated program length \t(Np):", h.Np)
	fmt.Println("Calculated program volume \t(V):", h.V)
	fmt.Println("Calculated program difficulty \t(D):", h.D)
	fmt.Println("Calculated program effort \t(E):", h.E)
	fmt.Println("Calculated program time \t(T):", h.T)
	fmt.Println("Calculated program bugs \t(B):", h.B)
	fmt.Println("------------------------------------------------------------")
}

type TokenType int

const (
	OPERATOR TokenType = iota
	OPERAND
)

type Token struct {
	// Token type
	Type TokenType
	// Token value
	Value string
}

func GetHalsteadMetrics(tokens []Token) Halstead {

	h := Halstead{}

	operators := make(map[string]bool)
	operands := make(map[string]bool)

	for _, token := range tokens {
		switch token.Type {
		case OPERATOR:
			operators[token.Value] = true
			h.n1++
		case OPERAND:
			operands[token.Value] = true
			h.n2++
		}
	}

	h.N1 = len(operators)
	h.N2 = len(operands)

	h.Calculate()

	return h
}

//cli input from args
func getCLIInput() string {
	var filepath string
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s -filepath <filepath>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&filepath, "filepath", "", "Filepath to the source code")
	flag.Parse()

	if filepath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return filepath
}

func main() {

	filepath := getCLIInput()

	metrics, err := analyzeSourceCode(filepath)
	if err != nil {
		panic(err)
	}

	metrics.Print()
}
