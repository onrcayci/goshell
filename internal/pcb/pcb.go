package pcb

type PCB struct {
	PC    int
	Start int
	End   int
	Next  *PCB
}

func New(start, end int) *PCB {
	newPCB := PCB{PC: start, Start: start, End: end}
	return &newPCB
}
