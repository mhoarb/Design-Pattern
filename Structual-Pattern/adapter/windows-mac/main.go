package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertlLightningUSB(mac)

	windowsMachin := &Windows{}
	WindowsMahcinAdapter := &WindowsAdapter{windowsMachine: windowsMachin}

	client.InsertlLightningUSB(WindowsMahcinAdapter)
}
