package main

type CounterTerrorist struct {
	color string
}

func (c *CounterTerrorist) getColor() string {
	return c.color
}

func newCounterTerroristDress() *TerroristDress {
	return &TerroristDress{color: "blue"}
}
