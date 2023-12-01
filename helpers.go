package compiler

func encodeFactName(fact string) []byte {
	// Example: encode the fact name as bytes
	return []byte(fact)
}

func encodeValue(value interface{}) []byte {
	// Example: encode values as bytes, handling different types
	// This is a simplified version. You'll need to handle each type properly.
	switch v := value.(type) {
	case int:
		return []byte{byte(v)}
	case string:
		return []byte(v)
		// Add cases for other types
	}
	return nil
}
