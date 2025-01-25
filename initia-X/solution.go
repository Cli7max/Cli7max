package main

import (
	"fmt"
	"cosmossdk.io/math"
	"cosmossdk.io/math/legacy"
	"github.com/cosmos/cosmos-sdk/types/math/legacy"
)
func (k TestAnteKeeper) GetBaseSpotPrice(denomQuote string) (math.LegacyDec, error) {
	balances, found := k.pools[denomQuote]
	if !found {
		return math.LegacyZeroDec(), fmt.Errorf("denomQuote %s not found in pools", denomQuote)
	}

	weights, found := k.weights[denomQuote]
	if !found {
		return math.LegacyZeroDec(), fmt.Errorf("denomQuote %s not found in weights", denomQuote)
	}

	// Fix: Check if balances and weights have at least 2 elements
	if len(balances) < 2 {
		return math.LegacyZeroDec(), fmt.Errorf("balances for denomQuote %s must have at least 2 elements", denomQuote)
	}

	if len(weights) < 2 {
		return math.LegacyZeroDec(), fmt.Errorf("weights for denomQuote %s must have at least 2 elements", denomQuote)
	}

	return types.GetBaseSpotPrice(balances[0], balances[1], weights[0], weights[1]), nil
}
