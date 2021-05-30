package memory

type mem struct {
	variable string
	value    string
}

var memory []*mem

func NewMemoryItem(variable, value string) *mem {
	memoryItem := mem{variable, value}
	return &memoryItem
}

func (m *mem) Set() {
	memory = append(memory, m)
}

func FindMemoryItem(variable string) string {
	for _, v := range memory {
		if v.variable == variable {
			return v.variable
		}
	}
	return ""
}
