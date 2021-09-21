package main

import (
	"fmt"
)

type astro struct {
	name string
	age int
	lastmission string
	isneeded bool
}

type nasaMission struct {
	people []astro
	number int
	message string
}

func main() {
	a1 := astro{
		name: "Bill",
		age: 50,
		lastmission: "long time ago",
		isneeded: false,
	}
        a2 := astro{
                name: "Al",
                age: 50,
                lastmission: "not long time ago",
                isneeded: false,
        }
        a3 := astro{
                name: "Ben",
                age: 50,
                lastmission: "long long time ago",
                isneeded: false,
        }
	x1 := []astro{a1, a2, a3}
	for _, v := range x1 {
		fmt.Println(v)
	}
	fmt.Println(x1)
	fmt.Println(x1[1].lastmission)

	s := nasaMission{
		people: x1,
		number: 3,
		message: "success",
	}
	fmt.Println(":",x1)
	fmt.Printf("::%+v", s)


}

