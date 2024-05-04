package main

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (oc *OnCommand) execute() {
	oc.device.on()
}

type OffCommand struct {
	device Device
}

func (oc *OffCommand) execute() {
	oc.device.off()
}
