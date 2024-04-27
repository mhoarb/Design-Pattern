package main

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}

type single struct {
	name string
}

var singleInstance *single

func getInstnce() *single {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now")
			singleInstance = &single{name: "mohammad"}
			fmt.Println(singleInstance)
		} else {
			fmt.Println("single instance already created")
		}
	} else {
		fmt.Println("single instance already created")
	}
	return singleInstance
}
