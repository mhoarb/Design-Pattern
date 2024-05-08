package main

func main() {
	lfu := &Lfu{}
	cash := initCash(lfu)

	cash.add("a", "1")
	cash.add("b", "2")
	cash.get("b")
	cash.add("c", "3")

	lru := &Lru{}
	cash2 := initCash(lru)
	cash2.add("a", "2")
	cash2.add("l", "2")
	cash2.add("p", "2")

}
