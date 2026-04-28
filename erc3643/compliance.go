package erc3643

import (
	"errors"

	"github.com/luxfi/geth/common"
)

// Compliance wraps the ModularCompliance contract attached to a Token.
type Compliance struct {
	Address common.Address
	client  Client
}

// NewCompliance binds to an existing ModularCompliance at addr.
func NewCompliance(client Client, addr common.Address) (*Compliance, error) {
	if client == nil {
		return nil, errors.New("erc3643: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("erc3643: zero address")
	}
	return &Compliance{Address: addr, client: client}, nil
}
