package main

type Image interface {
	Load(token string) (string, error)
	Display()
}
