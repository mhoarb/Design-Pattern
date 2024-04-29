package main

import "fmt"

func main() {
	nginxServer := NewNginxServer()
	appStatusUrl := "app/status"
	createUserURL := "create/user"

	response := nginxServer.HandleRequestNginx(appStatusUrl, "GET")
	fmt.Printf("\nurl: %s\nHttpStatusCode:%d\nBody: %s\n", appStatusUrl, response.Code, response.Msg)

	response = nginxServer.HandleRequestNginx(createUserURL, "GET")
	fmt.Printf("\nurl: %s\nHttpStatusCode:%d\nBody: %s\n", appStatusUrl, response.Code, response.Msg)

}
