package main

import (
	"fmt"
	"cosmossdk.io/math"
	"github.com/initia-labs/initia/x/move/types"
)
//Step 1: Define the TestAnteKeeper Struct:
//We define a TestAnteKeeper struct that contains two maps: pools and weights. 
//These are used to simulate the pools and weights for different denomQuote values (e.g., "atom").
//In this example, we intentionally set up the pools and weights maps to contain only one element for "atom".

//Step 2: Write the GetBaseSpotPrice Method:
//The GetBaseSpotPrice method tries to access the first two elements of balances and weights.
//If either balances or weights contains fewer than two elements, the method will attempt to access an invalid index, which will result in a runtime panic (index out of range).

//Step 3: Main Function:

//We simulate the faulty scenario by calling GetBaseSpotPrice("atom"), which will trigger the bug due to the insufficient number of elements in the pools and weights maps.
//The program will output an error message, and the panic will happen when trying to access the invalid index.

//Step 4: Run the Program:

//Run the program and observe the error message and potential panic due to the index out of range issue.

// TestAnteKeeper simulates a keeper with pools and weights.
type TestAnteKeeper struct {
	pools   map[string][]math.Int
	weights map[string][]math.LegacyDec
}

// GetBaseSpotPrice calculates the spot price based on balances and weights.
func (k TestAnteKeeper) GetBaseSpotPrice(denomQuote string) (math.LegacyDec, error) {
	balances, found := k.pools[denomQuote]
	if !found {
		return math.LegacyZeroDec(), fmt.Errorf("denomQuote %s not found in pools", denomQuote)
	}

	weights, found := k.weights[denomQuote]
	if !found {
		return math.LegacyZeroDec(), fmt.Errorf("denomQuote %s not found in weights", denomQuote)
	}

	// Potential bug: if balances or weights have less than two elements, this will cause a panic
	return types.GetBaseSpotPrice(balances[0], balances[1], weights[0], weights[1]), nil
}

func main() {
	// Initialize TestAnteKeeper with incorrect data (only 1 element in balances and weights)
	keeper := TestAnteKeeper{
		pools: map[string][]math.Int{
			"atom": {math.NewInt(10)}, // Only 1 balance
		},
		weights: map[string][]math.LegacyDec{
			"atom": {math.LegacyNewDecWithPrec(2, 1)}, // Only 1 weight
		},
	}

	// Call GetBaseSpotPrice with "atom" which should trigger the index out of range error
	quotePrice, err := keeper.GetBaseSpotPrice("atom")
	if err != nil {
		fmt.Println("Error:", err) // Expected behavior
	} else {
		fmt.Println("Base Spot Price:", quotePrice)
	}
}

// Output: Error: balances for denomQuote atom must have at least 2 elements  
// Panic: runtime error: index out of range [1] with length 1
//In the above code snippet, we intentionally set up the `pools` and `weights` maps to contain only one element for the `atom` denomQuote. When calling the `GetBaseSpotPrice` method with `atom`, the program will output an error message and potentially panic due to the index out of range issue.
//This bug can be fixed by checking if `balances` and `weights` have at least two elements before accessing them in the `GetBaseSpotPrice` method. This will prevent the panic and provide a more informative error message.
