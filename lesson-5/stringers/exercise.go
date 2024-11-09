package main

import "fmt"

// When fmt.Printf("%v: %v\n", name, ip) is executed, Go automatically inspects the ip value (of type IPAddr).
// If the type has a String() method defined, it is called, and the IP address is returned in string format.
type IPAddr [4]byte

// String() method dynamically constructs the string representation of the IP address.
func (ip IPAddr) String() string {
	var result string
	for _, b := range ip {
		if result != "" {
			result += "."
		}
		result += fmt.Sprintf("%d", b)
	}
	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
