package main

import "fmt"

func main() {
	newNginxServer := NewNginxServer()
	appStatusUrl := "/app/status"
	createUserUrl := "/create/user"

	responses := newNginxServer.HandleRequest(appStatusUrl, "GET")
	fmt.Printf("appStatusUrl: %s response: %d,%s\n", appStatusUrl, responses.Code, responses.Msg)

	responses = newNginxServer.HandleRequest(createUserUrl, "POST")
	fmt.Printf("appStatusUrl: %s response: %d,%s ", appStatusUrl, responses.Code, responses.Msg)
}
