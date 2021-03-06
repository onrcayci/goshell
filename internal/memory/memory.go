package memory

// Shell environment variable struct.
// Variable = name of the variable,
// Value = value the variable holds.
type mem struct {
	variable string
	value    string
}

// Shell environment variables slice.
// Holds the environment variables available in the current runtime.
var memory []*mem

// Function to add a new environment variable to the shell memory.
func NewMemoryItem(variable, value string) {
	for _, v := range memory {
		if v.variable == variable {
			v.value = value
			return
		}
	}
	memoryItem := &mem{variable, value}
	memory = append(memory, memoryItem)
}

// Function to find an environment variable with the given variable name.
func FindMemoryItem(variable string) string {
	for _, v := range memory {
		if v.variable == variable {
			return v.value
		}
	}
	return ""
}
