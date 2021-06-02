package cpu

type CPU struct {
	IP     int
	IR     string
	Quanta int
}

func New() *CPU {
	newCPU := CPU{}
	return &newCPU
}
