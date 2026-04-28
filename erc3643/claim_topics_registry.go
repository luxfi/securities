package erc3643

import (
	"errors"

	"github.com/luxfi/geth/common"
)

// ClaimTopicsRegistry wraps the registry of required claim topics for a token.
type ClaimTopicsRegistry struct {
	Address common.Address
	client  Client
}

// NewClaimTopicsRegistry binds to an existing ClaimTopicsRegistry.
func NewClaimTopicsRegistry(client Client, addr common.Address) (*ClaimTopicsRegistry, error) {
	if client == nil {
		return nil, errors.New("erc3643: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("erc3643: zero address")
	}
	return &ClaimTopicsRegistry{Address: addr, client: client}, nil
}
