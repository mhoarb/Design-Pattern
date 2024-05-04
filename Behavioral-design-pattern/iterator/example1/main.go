package main

import "fmt"

func main() {
	user1 := &User{
		name: "ALI",
		age:  18,
	}
	user2 := &User{
		name: "mmd",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %v\n", user)
	}
}
