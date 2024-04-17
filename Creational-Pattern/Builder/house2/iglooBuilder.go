package main

var _ Builder = &IglooBuilder{}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType(windowType string) {
	i.windowType = windowType
}

func (i *IglooBuilder) setDoorType(doorType string) {
	i.doorType = doorType
}

func (i *IglooBuilder) setNumFloor(floor int) {
	i.floor = floor
}
