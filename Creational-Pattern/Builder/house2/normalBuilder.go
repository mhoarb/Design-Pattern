package main

var _ Builder = &NormalBuilder{}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType(windowType string) {
	n.windowType = windowType
}

func (n *NormalBuilder) setDoorType(doorType string) {
	n.doorType = doorType
}

func (n *NormalBuilder) setNumFloor(floor int) {
	n.floor = floor
}
