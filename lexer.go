package main

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"

	LEFT_BRACKET  TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"

	STRING  TokenType = "STRING"
	BOOLEAN TokenType = "BOOLEAN"
	NUMBER  TokenType = "NUMBER"
	NULL    TokenType = "NULL"

	COLON TokenType = "COLON"
	COMMA TokenType = "COMMA"

	EOF     TokenType = "EOF"
	INVALID TokenType = "INVALID"
)

type Token struct {
	Type    TokenType
	Literal string
}

func lexer(input string) []Token {
	var tokens []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '{':
			tokens = append(tokens, Token{Type: LEFT_BRACE, Literal: "{"})
		case '}':
			tokens = append(tokens, Token{Type: RIGHT_BRACE, Literal: "}"})
		case '[':
			tokens = append(tokens, Token{Type: LEFT_BRACKET, Literal: "["})
		case ']':
			tokens = append(tokens, Token{Type: RIGHT_BRACKET, Literal: "]"})
		case ':':
			tokens = append(tokens, Token{Type: COLON, Literal: ":"})
		case ',':
			tokens = append(tokens, Token{Type: COMMA, Literal: ","})
		case '"':
			j := i + 1
			for j < len(input) && input[j] != '"' {
				j++
			}
			if j < len(input) {
				tokens = append(tokens, Token{Type: STRING, Literal: input[i : j+1]})
				i = j
			} else {
				tokens = append(tokens, Token{Type: INVALID, Literal: input[i:]})
				return tokens
			}
		default:
			if unicode.IsSpace(rune(input[i])) {
				continue
			} else if strings.HasPrefix(input[i:], "true") {
				tokens = append(tokens, Token{Type: BOOLEAN, Literal: "true"})
				i += 3
			} else if strings.HasPrefix(input[i:], "false") {
				tokens = append(tokens, Token{Type: BOOLEAN, Literal: "false"})
				i += 4
			} else if strings.HasPrefix(input[i:], "null") {
				tokens = append(tokens, Token{Type: NULL, Literal: "null"})
				i += 3
			} else if unicode.IsDigit(rune(input[i])) {
				j := i

				for j < len(input) && (unicode.IsDigit(rune(input[j])) || input[j] == '.') {
					j++
				}

				if input[i] == '0' {
					tokens = append(tokens, Token{Type: INVALID, Literal: string(input[i])})
				} else {
					tokens = append(tokens, Token{Type: NUMBER, Literal: input[i:j]})
				}

				i = j - 1
			} else {
				tokens = append(tokens, Token{Type: INVALID, Literal: string(input[i])})
				return tokens
			}
		}
	}

	tokens = append(tokens, Token{Type: EOF, Literal: ""})

	return tokens
}
