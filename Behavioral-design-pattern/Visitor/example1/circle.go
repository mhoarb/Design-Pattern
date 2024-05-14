package main 

type Circle struct {
	radius float64
}

func(c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

func(c *Circle) GetRadius() float64 {
	return c.radius
}

