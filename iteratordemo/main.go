package main

import (
	i "designpattern/iteratordemo/iterator"
	"fmt"
)

func main() {
	user1 := &i.User{
		Name: "Xuan Anh",
		Age:  18,
	}

	user2 := &i.User{
		Name: "Tom",
		Age:  40,
	}

	userCollection := &i.UserCollection{
		Users: []*i.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
