package main

import (
	"fmt"
)

func main() {
	const uri  = "https://example.org:6001/v2/snacks?"
	r := "req=snickers"
	q := "quantity=1"
 	s := "size=king"

	fmt.Printf("%v%v&%v&%v\n", uri, r, q, s)
}
