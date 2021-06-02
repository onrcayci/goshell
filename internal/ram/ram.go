package ram

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var RAM []string

func init() {
	RAM = make([]string, 1000)
}

func LoadToRAM(file *os.File) ([]int, error) {
	memoryAddresses := make([]int, 0)
	for index, value := range RAM {
		if value == "" {
			memoryAddresses = append(memoryAddresses, index)
			memoryAddresses = append(memoryAddresses, index)
			break
		}
	}
	if len(memoryAddresses) != 2 {
		return nil, errors.New("no more space in RAM")
	}
	reader := bufio.NewReader(file)
	var err error = nil
	line := ""
	for err != io.EOF {
		line, err = reader.ReadString('\n')
		if err != io.EOF && err != nil {
			return nil, err
		}
		RAM[memoryAddresses[1]] = line
		memoryAddresses[1]++
	}
	return memoryAddresses, nil
}

func FreeRAM(start, end int) {
	for i := start; i < end; i++ {
		RAM[i] = ""
	}
}
