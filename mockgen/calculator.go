package mockgen

//go:generate mockgen -source=calculator.go -destination=mock_calculator.go -package=mockgen
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
}

type Computer struct {
    calc Calculator
}

func NewComputer(calc Calculator) *Computer {
    return &Computer{calc: calc}
}

func (c *Computer) AddAndSubtract(a, b int) (int, int) {
    sum := c.calc.Add(a, b)
	// sum++ 
    difference := c.calc.Subtract(a, b)
    return sum, difference
}