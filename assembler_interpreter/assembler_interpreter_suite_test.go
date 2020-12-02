package assembler_interpreter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAssemblerInterpreter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AssemblerInterpreter Suite")
}
