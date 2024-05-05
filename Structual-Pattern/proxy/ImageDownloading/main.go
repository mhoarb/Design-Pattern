package main

import "fmt"

func main() {
	proxy := NewSecureProxyImage("https://buffer.com/library/content/images/2023/10/free-images.jpg")

	image, err := proxy.Load("555")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(image)

	proxy.Display()

	_, err = proxy.Load("Invalid_token")
	if err != nil {
		fmt.Println(err)
	}

}
