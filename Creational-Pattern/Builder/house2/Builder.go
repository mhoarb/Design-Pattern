package main

import "fmt"

type Builder interface {
	setWindowType(string)
	setDoorType(string)
	setNumFloor(int)
}

func getBuilder(builderType string) Builder {
	if builderType == "normal" {
		return newNormalBuilder()
	} else if builderType == "igloo" {
		return newIglooBuilder()
	}
	fmt.Println("your product not supported in my app")
	return nil

}
