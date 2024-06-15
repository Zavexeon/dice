package tokenizer

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	NumberToken TokenType = iota
	OperatorToken
	LeftParenthesisToken
	RightParenthesisToken
)

type Token struct {
	Type    TokenType
	Content string
}

func Tokenize(expression string) ([]Token, error) {
	tokenizedExpression := []Token{}
	numberBuffer := []string{}

	for _, characterRune := range expression {
		if unicode.IsSpace(characterRune) {
			continue
		}

		if unicode.IsDigit(characterRune) {
			numberBuffer = append(numberBuffer, string(characterRune))
			continue
		}

		switch characterRune {
		case '+', '-', 'x', '/', '^':
			if len(numberBuffer) > 0 {
				tokenizedExpression = append(tokenizedExpression, Token{NumberToken, strings.Join(numberBuffer, "")})
				numberBuffer = []string{}
			}
			tokenizedExpression = append(tokenizedExpression, Token{OperatorToken, string(characterRune)})
		case '(':
			tokenizedExpression = append(tokenizedExpression, Token{LeftParenthesisToken, "("})
		case ')':
			tokenizedExpression = append(tokenizedExpression, Token{RightParenthesisToken, ")"})
		case '.':
			if len(numberBuffer) == 0 {
				numberBuffer = append(numberBuffer, "0.")
			}
		default:
			return nil, fmt.Errorf("invalid token in expression: %c", characterRune)
		}
	}

	if len(numberBuffer) > 0 {
		tokenizedExpression = append(tokenizedExpression, Token{NumberToken, strings.Join(numberBuffer, "")})
	}

	return tokenizedExpression, nil
}
