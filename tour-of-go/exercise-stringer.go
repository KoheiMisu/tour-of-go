package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// add a "String()" method to IPAddr

func (i IPAddr) String() string {
	var joined string

	for index, block := range i {
		if index != 3 {
			// blockはuint8なので下記のようにstringにする
			val := strconv.Itoa(int(block)) + "."
			joined += val
			continue
		}
		joined += strconv.Itoa(int(block))
	}

	return joined
}

func main()  {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
