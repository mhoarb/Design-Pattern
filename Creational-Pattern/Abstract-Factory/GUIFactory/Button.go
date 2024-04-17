package main

type Button interface {
	Click()
}

var _ Button = WindowsButton{}

type WindowsButton struct{}

func (w WindowsButton) Click() {
	//TODO implement me
}

var _ Button = &MacButton{}

type MacButton struct{}

func (m MacButton) Click() {
	//TODO implement me
}
