//Expected query for address: cosmos1exampleaccount
//Actual query made for address: cosmos1hardcodedaddress

package main

import (
	"fmt"
	"log"
)

// Mock implementation for demonstration
func TestHardcodedAddressBug(addr string, hardcodedAddress string) {
	if addr != hardcodedAddress {
		log.Printf("Bug Detected: Expected query for address %s, but queried hardcoded address %s", addr, hardcodedAddress)
	}
}

func main() {
	TestHardcodedAddressBug("cosmos1exampleaccount", "cosmos1hardcodedaddress")
}
// Bug Detected: Expected query for address cosmos1exampleaccount, but queried hardcoded address cosmos1hardcodedaddress

