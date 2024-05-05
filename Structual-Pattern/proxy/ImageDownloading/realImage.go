package main

import "fmt"

var _ Image = &RealIMage{}

type RealIMage struct {
	url string
}

func (r *RealIMage) Load(token string) (string, error) {
	if token != "555" {
		return "", fmt.Errorf("Invalid access token")
	}
	fmt.Println("Downloading image from", r.url)
	return "Image Dwnloading", nil
}

func (r *RealIMage) Display() {
	fmt.Println("displaying normal image")
}
