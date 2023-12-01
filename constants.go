package compiler

const (
	OpEqual = iota
	OpGreaterThan
	// ... other comparison operations

	OpAnd
	OpOr
	// ... other logical operations

	OpLoadFact
	OpStoreFact

	OpJumpIfTrue
	OpJumpIfFalse

	OpExecuteAction
)
