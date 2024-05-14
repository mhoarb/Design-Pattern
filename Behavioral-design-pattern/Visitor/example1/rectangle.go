package main 

type Rectangle struct {
	length float64
	width float64
}

func(r *Rectangle)Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}


func(r *Rectangle) GetLength() float64 {
	return r.length
}

func(r *Rectangle)GetSide() float64 {
	return r.width
}