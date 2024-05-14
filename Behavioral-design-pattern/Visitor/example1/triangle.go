package main

type Triangle struct {
	base float64
	height float64
	hypotenuse float64
}

func(t *Triangle) Accept(visitor Visitor) {
	visitor.VisitTriangle(t)
}

func(t *Triangle) GetBase()float64 {
	return t.base
}

func(t *Triangle) GetHeight() float64 {
	return t.height
}

func(t *Triangle)GetHypotenuse() float64 {
	return t.hypotenuse
}

