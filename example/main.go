package main

import (
	"context"

	ga "github.com/ActiveState/go-ogle-analytics"
)

func main() {
	client, err := ga.NewClient("UA-30305960-4")
	if err != nil {
		panic(err)
	}

	err = client.SendWithContext(context.Background(), ga.NewEvent("Foo", "Bar").Label("Bazz"))
	if err != nil {
		panic(err)
	}

	println("Event fired!")
}
