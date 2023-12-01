package compiler

import (
	"fmt"

	"github.com/rgehrsitz/rulegopher/pkg/rules"
)

// Assuming a struct for Actions exists
func compileAction(action rules.Action) ([]byte, error) {
	var bytecode []byte

	// Example: Convert action to bytecode based on its type
	switch action.Type {
	case "Notify":
		bytecode = append(bytecode, OpExecuteAction)
		// Add additional operands as required
	// Handle other action types
	default:
		return nil, fmt.Errorf("unsupported action type: %s", action.Type)
	}

	return bytecode, nil
}
