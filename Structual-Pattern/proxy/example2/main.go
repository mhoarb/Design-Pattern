package main

func main() {
	processor := &RealImageProcessor{}
	proxy := NewCachingProxy(processor)
	img, err := LoadImage("path/to/image.jpg")

	err = proxy.Resize(img, 100, 100)
	if err != nil {
		return
	}
}
