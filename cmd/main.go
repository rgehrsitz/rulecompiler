package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rgehrsitz/rulegopher/pkg/rules"
)

// readAndParseRules reads a JSON file containing rules and parses it into a slice of Rule structs.
func readAndParseRules(filePath string) ([]rules.Rule, error) {
	var ruleSet []rules.Rule

	// Read the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to Rule structs
	err = json.Unmarshal(fileData, &ruleSet)
	if err != nil {
		return nil, err
	}

	return ruleSet, nil
}

func main() {
	// Example usage
	ruleFilePath := "C:/code/rulegopher/pkg/engine/rules.json"
	rules, err := readAndParseRules(ruleFilePath)
	if err != nil {
		log.Fatalf("Failed to read or parse rules: %v", err)
	}

	log.Printf("Successfully read and parsed %d rules", len(rules))
	// Enhanced AddRule or a similar compiler function
	for _, rule := range rules {
		bytecode, err := compiler.CompileRule(rule)
		if err != nil {
			log.Printf("Failed to compile rule %s: %v", rule.Name, err)
			continue
		}

		// Store or use bytecode...
	}
}
