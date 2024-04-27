package main

import "fmt"

type Clinet struct {
	name string
}

func (c *Clinet) Insert(computer Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	computer.InsertIntoLightningPort()
}
