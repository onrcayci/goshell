package kernel

import (
	"container/list"
	"os"

	"github.com/onrcayci/goshell/internal/cpu"
	"github.com/onrcayci/goshell/internal/pcb"
	"github.com/onrcayci/goshell/internal/ram"
)

var ReadyQueue *list.List
var RuntimeCPU *cpu.CPU

func init() {
	ReadyQueue = list.New()
	RuntimeCPU = cpu.New()
}

func MyInit(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	memoryAddresses, err := ram.LoadToRAM(file)
	if err != nil {
		return err
	}
	newPCB := pcb.New(memoryAddresses[0], memoryAddresses[1])
	ReadyQueue.PushBack(newPCB)
	return nil
}
