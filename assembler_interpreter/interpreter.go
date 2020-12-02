package assembler_interpreter

func SimpleAssembler(program []string) map[string]int {
	registers := make(map[string]int)

	for i := 0; i < len(program); i++ {
		tokens := Lexer(program[i])

		if tokenTypesMatch(tokens, []Type{MOV, REG, INT}) {
			registers[tokens[1].Value.register()] = tokens[2].Value.value()
		}
		if tokenTypesMatch(tokens, []Type{MOV, REG, REG}) {
			registers[tokens[1].Value.register()] = registers[tokens[2].Value.register()]
		}
		if tokenTypesMatch(tokens, []Type{DEC, REG}) {
			registers[tokens[1].Value.register()] = registers[tokens[1].Value.register()] - 1
		}
		if tokenTypesMatch(tokens, []Type{INC, REG}) {
			registers[tokens[1].Value.register()] = registers[tokens[1].Value.register()] + 1
		}
		if tokenTypesMatch(tokens, []Type{JNZ, REG, INT}) {
			if registers[tokens[1].Value.register()] != 0 {
				i += tokens[2].Value.value() - 1
			}
		}
		if tokenTypesMatch(tokens, []Type{JNZ, INT, INT}) {
			if tokens[1].Value.value() != 0 {
				i += tokens[2].Value.value() - 1
			}
		}
	}
	return registers
}

func tokenTypesMatch(tokens []Token, types []Type) bool {
	for i, tokenType := range types {
		if tokens[i].Type != tokenType {
			return false
		}
	}

	return true
}
