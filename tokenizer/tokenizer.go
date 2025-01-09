package tokenizer

import (
	"fmt"
	"halstead/colors"
	"os"
	"regexp"
	"strings"
)

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

func createLexer(filename string) *Lexer {

	// Read the source code from the file
	sourceCode, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lang := filename[strings.LastIndex(filename, ".")+1:]

	colors.GREEN.Printf("Language: %s\n", lang)

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