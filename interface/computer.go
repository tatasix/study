package _interface

type ComputerCpuInterface interface {
	Calculate(string)
}
type ComputerCardInterface interface {
	Display(string)
}
type ComputerMemoryInterface interface {
	Storage(string)
}

type Computer struct {
	cpu    ComputerCpuInterface
	card   ComputerCardInterface
	memory ComputerMemoryInterface
}

func NewComputer(cpu ComputerCpuInterface, memory ComputerMemoryInterface, card ComputerCardInterface) *Computer {
	return &Computer{
		cpu:    cpu,
		memory: memory,
		card:   card,
	}
}

func (c *Computer) DoWork() {
	c.cpu.Calculate("Computer cpu Calculate")
	c.memory.Storage("Computer memory Storage")
	c.card.Display("Computer card Display")
}
