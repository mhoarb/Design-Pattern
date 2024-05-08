package main

type EvictionAlgo interface {
	evict(c *Cash)
}
