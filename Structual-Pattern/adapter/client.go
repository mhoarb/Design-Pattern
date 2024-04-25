package main

import "fmt"

type Client struct{}

func (c *Client) InsertlLightningUSB(com Computer) {
	fmt.Println("client insert the lightning usb into computer")
	com.insertInToLightningPort()
}
