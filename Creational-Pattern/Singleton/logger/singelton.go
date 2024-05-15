package main

import (
	"fmt"
	"sync"

)

type MyLogger struct {

}



func GetMyLoggerInstance(i int) *MyLogger{
	var globalLogger *MyLogger
	mu := &sync.Mutex{}


	if globalLogger == nil {
		mu.Lock()
		defer mu.Unlock()
		
		fmt.Println("creating my logger instance by" , i)
		globalLogger = &MyLogger{}
	}else {
		fmt.Println("myLogger instance already created")
	}
	return globalLogger
}