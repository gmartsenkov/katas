package assembler_interpreter

import (
	"regexp"
	"strconv"
	"strings"
)

type TokenValue interface {
	value() int
	register() string
}

type TokenRegister struct {
	val string
}

func (tr TokenRegister) register() string {
	return tr.val
}

func (tr TokenRegister) value() int {
	return 0
}

type TokenInt struct {
	val int
}

func (ti TokenInt) register() string {
	return ""
}

func (ti TokenInt) value() int {
	return ti.val
}

type Type int

const (
	MOV = iota // Move value/register content to a register "mov x 5" or "mov x a"
	INC // Increment register "INC A"
	DEC // Decrement register "DEC B"
	JNZ // Move instruction relative to itself
	INT // Integer
	REG // Register
)

type Token struct {
	Type Type
	Value TokenValue
}

var IsLetter = regexp.MustCompile(`^[a-z]+$`).MatchString

func Lexer(command string) []Token {
	tokens := []Token{}

	for _, token := range strings.Split(command, " ") {
		switch token {
		case "mov":
			tokens = append(tokens, Token{Type: MOV})
		case "inc":
			tokens = append(tokens, Token{Type: INC})
		case "dec":
			tokens = append(tokens, Token{Type: DEC})
		case "jnz":
			tokens = append(tokens, Token{Type: JNZ})
		default:
			if IsLetter(token) {
				tokens = append(tokens, Token{Type: REG, Value: TokenRegister{val: token}})
			}

			if i, err := strconv.Atoi(token); err == nil {
				tokens = append(tokens, Token{Type: INT, Value: TokenInt{val: i}})
			}
		}
	}

	return tokens
}
