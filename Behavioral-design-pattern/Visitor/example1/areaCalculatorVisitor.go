package main

import "math"

type AreaCalculatorVisitor struct {
	totalArea float64
}

func(av *AreaCalculatorVisitor) VisitSquare(s *Square) {
	av.totalArea = s.side * s.side
}

func(av *AreaCalculatorVisitor) VisitCircle(c *Circle) {
	av.totalArea = math.Pi * c.radius * c.radius
}

func(av *AreaCalculatorVisitor) VisitTriangle(t *Triangle) {
	av.totalArea = (t.base * t.height) / 2
}

func(av *AreaCalculatorVisitor) VisitRectangle(r *Rectangle) {
	av.totalArea = r.length * r.width
}

