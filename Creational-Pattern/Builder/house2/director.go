package main

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{builder: b}
}
func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse(builder NormalBuilder) *House {
	d.builder.setWindowType(builder.windowType)
	d.builder.setDoorType(builder.doorType)
	d.builder.setNumFloor(builder.floor)
	return &House{}

}
