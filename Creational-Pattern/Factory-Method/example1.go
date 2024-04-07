package main

type Shape interface {
	Draw() string
}

type ShapeFactory interface {
	CreateShape(shapeType string) Shape
}
type Circle struct{}

func (c *Circle) Draw() string {
	return "0"
}

type Square struct{}

func (s *Square) Draw() string {
	return "-----\n|   |\n|   |n-----\n"
}
