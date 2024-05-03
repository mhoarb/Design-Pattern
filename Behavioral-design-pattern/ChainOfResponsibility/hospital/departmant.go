package main

type Departmant interface {
	execute(*Patient)
	setNext(Departmant)
}
