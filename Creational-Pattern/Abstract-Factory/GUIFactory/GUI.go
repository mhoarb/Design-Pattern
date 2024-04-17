package main

/*
*
its abstract factory
*/
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}
