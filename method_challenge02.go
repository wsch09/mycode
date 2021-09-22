/*
Create a strut type, Virtmach, to track information about virtual machines. 
Your record can track whatever you'd like, but values like ip, hostname, diskgb, ram, could all be possible values. 
Create at least two (2) methods that allow you to interact with your strut. If you come up with your own working solution, push it to your SCM (GitHub)
, then share a link to it with the class (if you're in an online class, post in the chat).
*/

package main

import (
	"fmt"
)

type vm struct {
	ip string
	hostname string
	diskgb int
	ram int
}

func (v *vm) newram(newRam int) {
	v.ram = newRam
}

func vmreport(v *vm) {
	fmt.Printf("laptop %v has %v gb of ram\n", v.hostname, v.ram)
}

func main() {
	virtmachine := vm {
		ip: "10.0.0.999",
		hostname: "bills laptop",
		diskgb: 500,
		ram: 256,
	}
	vmreport(&virtmachine)
	virtmachine.newram(512)
	vmreport(&virtmachine)

}

