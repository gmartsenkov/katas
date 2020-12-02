package assembler_interpreter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Tests", func() {
	It("should work for some basic example programs", func() {
		Expect(SimpleAssembler([]string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"})).To(Equal(map[string]int{"a": 1}))
		Expect(SimpleAssembler([]string{"mov a -10", "mov b a", "inc a", "dec b", "jnz a -2"})).To(Equal(map[string]int{"a": 0, "b": -20}))
	})

	Describe("Lexer", func() {
		It("parses the tokens correctly", func() {
			Expect(Lexer("mov a 5")).To(Equal([]Token{
				{Type: MOV},
				{Type: REG, Value: TokenRegister{val: "a"}},
				{Type: INT, Value: TokenInt{val: 5}},
			}))
			Expect(Lexer("dec a")).To(Equal([]Token{
				{Type: DEC},
				{Type: REG, Value: TokenRegister{val: "a"}},
			}))
			Expect(Lexer("inc b")).To(Equal([]Token{
				{Type: INC},
				{Type: REG, Value: TokenRegister{val: "b"}},
			}))
			Expect(Lexer("jnz -1")).To(Equal([]Token{
				{Type: JNZ},
				{Type: INT, Value: TokenInt{val: -1}},
			}))
		})
	})
	Describe("tokensMatch", func() {

	})
})
