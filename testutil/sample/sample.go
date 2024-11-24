package sample

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func Coin() *sdk.Coin {
	coin := sdk.NewCoin("uOPT", math.NewInt(100))
	return &coin
}
