package erc3643

import (
	"context"
	"errors"

	"github.com/luxfi/geth/common"
)

// Client is the minimal ETH-client surface this package needs.
type Client interface {
	ChainID(ctx context.Context) (any, error)
}

// Token wraps the ERC-3643 Token contract.
type Token struct {
	Address common.Address
	client  Client
}

// NewToken binds to an existing ERC-3643 Token at addr.
func NewToken(client Client, addr common.Address) (*Token, error) {
	if client == nil {
		return nil, errors.New("erc3643: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("erc3643: zero address")
	}
	return &Token{Address: addr, client: client}, nil
}
