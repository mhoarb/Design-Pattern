package main

import "time"

func main() {

	for i := 0; i <= 5; i++ {
		go GetMyLoggerInstance(i)
	}
	time.Sleep(time.Second)
}
