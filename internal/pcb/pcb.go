package pcb

type PCB struct {
	IP    int
	Start int
	End   int
	Next  *PCB
}

func New(start, end int) *PCB {
	newPCB := PCB{IP: start, Start: start, End: end}
	return &newPCB
}
