package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// Recommended Mitigation Steps - Add Input Validation:
// Ensure host and challenger are validated before assignment

func validateParticipants(host, challenger interface{}) error {
    if host == nil {
        return errors.New("host cannot be nil")
    }
    if challenger == nil {
        return errors.New("challenger cannot be nil")
    }
    return nil
}

func main() {
    // Example usage
    host := "HostName"
    challenger := "ChallengerName"
    
    err := validateParticipants(host, challenger)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Both participants are valid.")
    }
}

//Test Edge Cases:
//Write unit tests to validate that the function handles nil inputs appropriately:

type Child struct{}

func (c *Child) Initialize(host, challenger interface{}) error {
    return validateParticipants(host, challenger)
}

func TestInitialize_NilInputs(t *testing.T) {
    var nilHost interface{}
    var nilChallenger interface{}

    ch := &Child{}
    err := ch.Initialize(nilHost, nilChallenger)
    if err == nil || !strings.Contains(err.Error(), "host cannot be nil") {
        t.Fatalf("expected error for nil host, got: %v", err)
    }
}


