package onchainid

import (
	"errors"

	"github.com/luxfi/geth/common"
)

// IdentityFactory wraps the IdFactory upstream contract used to deploy
// per-investor ONCHAINID Identities.
type IdentityFactory struct {
	Address common.Address
	client  Client
}

// NewIdentityFactory binds to an existing IdFactory at addr.
func NewIdentityFactory(client Client, addr common.Address) (*IdentityFactory, error) {
	if client == nil {
		return nil, errors.New("onchainid: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("onchainid: zero address")
	}
	return &IdentityFactory{Address: addr, client: client}, nil
}
