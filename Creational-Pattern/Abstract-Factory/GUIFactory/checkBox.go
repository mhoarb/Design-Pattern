package main

type Checkbox interface {
	Check() bool
	UnCheck() bool
}

var _ Checkbox = &WindowsCheckbox{}

type WindowsCheckbox struct {
	checked bool
}

func (w *WindowsCheckbox) Check() bool {
	w.checked = true
	return w.checked
}
func (w *WindowsCheckbox) UnCheck() bool {
	w.checked = false
	return w.checked
}

var _ Checkbox = &MacCheckbox{}

type MacCheckbox struct {
	checked bool
}

func (m *MacCheckbox) Check() bool {
	m.checked = true
	return m.checked
}
func (m *MacCheckbox) UnCheck() bool {
	m.checked = false
	return m.checked

}
