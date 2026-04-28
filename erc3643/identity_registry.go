package erc3643

import (
	"errors"

	"github.com/luxfi/geth/common"
)

// IdentityRegistry wraps the per-token IdentityRegistry.
type IdentityRegistry struct {
	Address common.Address
	client  Client
}

// NewIdentityRegistry binds to an existing IdentityRegistry at addr.
func NewIdentityRegistry(client Client, addr common.Address) (*IdentityRegistry, error) {
	if client == nil {
		return nil, errors.New("erc3643: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("erc3643: zero address")
	}
	return &IdentityRegistry{Address: addr, client: client}, nil
}
