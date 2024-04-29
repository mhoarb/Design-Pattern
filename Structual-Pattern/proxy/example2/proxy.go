package main

import (
	"fmt"
)

type CachingProxy struct {
	realProcessor ImageProcessor
	cache         map[string]string
}

func NewCachingProxy(processor ImageProcessor) *CachingProxy {
	return &CachingProxy{
		realProcessor: processor,
		cache:         make(map[string]string),
	}
}

func (c *CachingProxy) Resize(img string, width, height int) error {
	key := fmt.Sprintf("resized-%d * %d", width, height)
	if cachedImg, ok := c.cache[key]; ok {
		fmt.Println(cachedImg)
		return nil
	}
	err := c.realProcessor.Resize(img, width, height)
	if err != nil {
		return err
	}
	c.cache[key] = img
	return nil
}
