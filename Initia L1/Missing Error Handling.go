package main

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic")
		}
	}()

	encodingConfig := codec.MakeCodec()
	require.NotNil(t, encodingConfig, "encodingConfig should not be nil")

	// Test RegisterConcrete with nil values
	require.NotPanics(t, func() {
		encodingConfig.Amino.RegisterConcrete(nil, "testdata.TestMsg", nil)
	}, "RegisterConcrete should not panic with nil values")
}