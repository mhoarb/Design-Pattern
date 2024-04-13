package main

type Document interface {
	GetFormat() string
	Clone() Document
}
