package main

func main() {
	square := Square{side: 5.0}
	rectangle := Rectangle{length: 3.0 , width: 4.0}
	circle := Circle{radius: 6.0}
	triangle := Triangle{base: 7.0 , height: 8.0 , hypotenuse:9.0 }

	areaCalculator := AreaCalculatorVisitor{}
	
	square.Accept(&areaCalculator)
	rectangle.Accept(&areaCalculator)
	circle.Accept(&areaCalculator)
	triangle.Accept(&areaCalculator)


	perimeterCalculatorVisitor := PerimeterCalculatorVisitor{}

	square.Accept(&perimeterCalculatorVisitor)
	rectangle.Accept(&perimeterCalculatorVisitor)
	circle.Accept(&perimeterCalculatorVisitor)
	triangle.Accept(&perimeterCalculatorVisitor)
	
}