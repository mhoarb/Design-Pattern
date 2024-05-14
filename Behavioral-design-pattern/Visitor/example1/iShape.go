package main

type Shape interface {
	Accept(Visitor)
}