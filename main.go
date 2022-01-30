package main

import (
	"fmt"

	"github.com/kotaoue/go-wd"
)

func main() {
	dir, err := wd.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(dir)
}
