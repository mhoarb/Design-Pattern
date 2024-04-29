package main

type ImageProcessor interface {
	Resize(img string, width, height int) error
}

func LoadImage(img string) (string, error) {
	return img, nil
}
