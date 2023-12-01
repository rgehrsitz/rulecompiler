package compiler

import (
	"reflect"
	"testing"

	"github.com/rgehrsitz/rulegopher/pkg/rules"
)

func TestCompileCondition(t *testing.T) {
	// Test for a simple equality condition
	cond := rules.Condition{Fact: "testFact", Operator: "equal", Value: 10}
	expected := []byte{OpLoadFact}
	expected = append(expected, encodeFactName("testFact")...)
	expected = append(expected, OpEqual)
	expected = append(expected, encodeValue(10)...)

	result, err := compileCondition(cond)
	if err != nil {
		t.Fatalf("compileCondition failed: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Additional tests for other operators and complex conditions
}
