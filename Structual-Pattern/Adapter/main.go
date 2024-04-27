package main

func main() {
	client := &Clinet{}
	mac := &Mac{}

	client.Insert(mac)

	windowsMachin := &Windows{}
	windowsMAchinAdapter := &WindowsAdapter{windwsMachin: windowsMachin}

	client.Insert(windowsMAchinAdapter)
}
