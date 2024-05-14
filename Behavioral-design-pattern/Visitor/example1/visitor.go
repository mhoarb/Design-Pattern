package main 


type Visitor interface {
	VisitSquare(s *Square)
	VisitRectangle(r *Rectangle)
	VisitCircle(c *Circle)
	VisitTriangle(c *Triangle)
	
}




















