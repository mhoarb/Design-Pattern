package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	director := newDirector(normalBuilder)
	normalBuilderDetail := NormalBuilder{
		windowType: "wooden",
		doorType:   "alp",
		floor:      1,
	}

	normalHouse := director.buildHouse(normalBuilderDetail)
	fmt.Printf("normal house WindowType:%v\n, doorTypr:%v\n, floor:%v\n",
		normalBuilderDetail.windowType, normalBuilderDetail.doorType, normalBuilderDetail.floor)
	fmt.Println(normalHouse.doorType)

}
