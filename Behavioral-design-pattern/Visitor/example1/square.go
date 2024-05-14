package main 


type Square struct {
	side float64
}


func(s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}

func(s *Square)GetSide() float64 {
	return s.side
}

