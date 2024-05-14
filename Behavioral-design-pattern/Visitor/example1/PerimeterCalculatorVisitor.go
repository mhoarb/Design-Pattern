package main

import "math"

type PerimeterCalculatorVisitor struct {
	totalPerimeter float64
}

func(pv *PerimeterCalculatorVisitor) VisitSquare(s *Square) {
	pv.totalPerimeter = 4 *s.side
}

func(pv *PerimeterCalculatorVisitor) VisitCircle(c *Circle) {
	pv.totalPerimeter = 2 * math.Phi *c.radius
}

func(pv *PerimeterCalculatorVisitor) VisitTriangle(t *Triangle) {
	pv.totalPerimeter = t.base + t.height + t.hypotenuse
}

func(pv *PerimeterCalculatorVisitor) VisitRectangle(r *Rectangle) {
	pv.totalPerimeter = (r.length + r.width) *2
}
