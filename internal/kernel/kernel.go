package kernel

import (
	"os"

	"github.com/onrcayci/goshell/internal/cpu"
	"github.com/onrcayci/goshell/internal/pcb"
	"github.com/onrcayci/goshell/internal/queue"
	"github.com/onrcayci/goshell/internal/ram"
)

var ReadyQueue *queue.Queue
var RuntimeCPU *cpu.CPU

func init() {
	ReadyQueue = queue.New()
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
	ReadyQueue.Push(newPCB)
	return nil
}

func Scheduler() {
	for ReadyQueue.Head != nil {
		if RuntimeCPU.Quanta == 0 || RuntimeCPU.IR == "" {
			currentPCB := ReadyQueue.Pop()
			RuntimeCPU.IP = currentPCB.PC
			quanta := currentPCB.End - currentPCB.PC
			if quanta == 0 {
				ram.FreeRAM(currentPCB.Start, currentPCB.End)
				continue
			} else if quanta > 2 {
				quanta = 2
			}
			RuntimeCPU.Quanta = quanta
			RuntimeCPU.Run()
			currentPCB.PC = RuntimeCPU.IP
			ReadyQueue.Push(currentPCB)
		}
	}
}
