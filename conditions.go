package compiler

import (
	"fmt"

	"github.com/rgehrsitz/rulegopher/pkg/rules"
)

func compileCondition(cond rules.Condition) ([]byte, error) {
	var bytecode []byte

	bytecode = append(bytecode, OpLoadFact)
	bytecode = append(bytecode, encodeFactName(cond.Fact)...)

	switch cond.Operator {
	case "equal":
		bytecode = append(bytecode, OpEqual)
		bytecode = append(bytecode, encodeValue(cond.Value)...)
	case "greaterThan":
		bytecode = append(bytecode, OpGreaterThan)
		bytecode = append(bytecode, encodeValue(cond.Value)...)
	// Add cases for other operators
	default:
		return nil, fmt.Errorf("unsupported operator: %s", cond.Operator)
	}

	return bytecode, nil
}
