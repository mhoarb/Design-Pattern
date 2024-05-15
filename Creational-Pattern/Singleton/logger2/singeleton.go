package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {

}

func GetMyLoggerInstance(i int) *MyLogger {
	var (
		globalLogger *MyLogger
		once sync.Once
	)

	if globalLogger == nil {
		once.Do(
			func() {
				fmt.Println("creating myLogger instance by" , i)
			})
	}else {
		fmt.Println("myLogger instance already created")
	}
	return globalLogger
}