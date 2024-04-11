package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}
func (i *iglooBuilder) setWindowType() {
	i.windowType = "Snow window"
}
func (i *iglooBuilder) setDoorType() {
	i.doorType = "Snow door"
}
func (i *iglooBuilder) setNumFloor() {
	i.floor = 1
}
func (i *iglooBuilder) getHouse() House {
	return House{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}
