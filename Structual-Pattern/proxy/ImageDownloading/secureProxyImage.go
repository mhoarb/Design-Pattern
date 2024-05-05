package main

import (
	"fmt"
	"sync"
)

type SecureProxyImage struct {
	realImage Image
	cache     map[string]string
	mutex     sync.Mutex
}

func NewSecureProxyImage(url string) *SecureProxyImage {
	return &SecureProxyImage{
		realImage: &RealIMage{url: url},
		cache:     make(map[string]string),
	}
}

func (s *SecureProxyImage) Load(token string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if image, ok := s.cache[token]; ok {
		fmt.Println("Image loaded from cache")
		return image, nil
	}
	image, err := s.realImage.Load(token)
	if err != nil {
		return "", nil
	}
	s.cache[token] = image
	return image, nil
}

func (s *SecureProxyImage) Display() {
	s.realImage.Display()
}
