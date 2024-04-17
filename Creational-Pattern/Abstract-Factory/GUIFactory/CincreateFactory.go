package main

var _ GUIFactory = &WindowsGui{}

type WindowsGui struct{}

func (w *WindowsGui) CreateButton() Button {
	return &WindowsButton{}
}

func (w *WindowsGui) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

var _ GUIFactory = &MacGUI{}

type MacGUI struct{}

func (m *MacGUI) CreateButton() Button {
	return &MacButton{}
}

func (m *MacGUI) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}
