package main

import (
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func main() {
    encodingConfig := codec.MakeTestEncodingConfig()
    err := encodingConfig.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)
    if err != nil {
        suite.T().Fatal("Failed to register TestMsg:", err)
    }
}
