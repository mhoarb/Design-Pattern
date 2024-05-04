package main

var _ Collection = &UserCollection{}
type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{}
}

