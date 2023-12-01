package compiler

import "github.com/rgehrsitz/rulegopher/pkg/rules"

func CompileRule(rule rules.Rule) ([]byte, error) {
	var bytecode []byte

	// Compile each condition in the rule
	for _, cond := range rule.Conditions.All {
		condBytecode, err := compileCondition(cond)
		if err != nil {
			return nil, err
		}
		bytecode = append(bytecode, condBytecode...)
	}

	// Compile events if present in the rule

	return bytecode, nil
}
