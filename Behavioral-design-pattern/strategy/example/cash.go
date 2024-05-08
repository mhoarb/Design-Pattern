package main

type Cash struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCash(e EvictionAlgo) *Cash {
	storage := make(map[string]string)
	return &Cash{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}
func (c *Cash) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}
func (c *Cash) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}
func (c *Cash) get(key string) {
	delete(c.storage, key)
}

func (c *Cash) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
