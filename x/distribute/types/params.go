package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxSupply = []byte("MaxSupply")
	// TODO: Determine the default value
	DefaultMaxSupply uint64 = 0
)

var (
	KeyAuthorizedAccounts = []byte("AuthorizedAccounts")
	// TODO: Determine the default value
	DefaultAuthorizedAccounts []string = []string{}
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxSupply uint64,
	authorizedAccounts []string,
) Params {
	return Params{
		MaxSupply:          maxSupply,
		AuthorizedAccounts: authorizedAccounts,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxSupply,
		DefaultAuthorizedAccounts,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxSupply, &p.MaxSupply, validateMaxSupply),
		paramtypes.NewParamSetPair(KeyAuthorizedAccounts, &p.AuthorizedAccounts, validateAuthorizedAccounts),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxSupply(p.MaxSupply); err != nil {
		return err
	}

	if err := validateAuthorizedAccounts(p.AuthorizedAccounts); err != nil {
		return err
	}

	return nil
}

// validateMaxSupply validates the MaxSupply param
func validateMaxSupply(v interface{}) error {
	maxSupply, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxSupply

	return nil
}

// validateAuthorizedAccounts validates the AuthorizedAccounts param
func validateAuthorizedAccounts(v interface{}) error {
	authorizedAccounts, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = authorizedAccounts

	return nil
}
