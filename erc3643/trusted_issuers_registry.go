package erc3643

import (
	"errors"

	"github.com/luxfi/geth/common"
)

// TrustedIssuersRegistry wraps the registry of approved claim issuers.
type TrustedIssuersRegistry struct {
	Address common.Address
	client  Client
}

// NewTrustedIssuersRegistry binds to an existing TrustedIssuersRegistry.
func NewTrustedIssuersRegistry(client Client, addr common.Address) (*TrustedIssuersRegistry, error) {
	if client == nil {
		return nil, errors.New("erc3643: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("erc3643: zero address")
	}
	return &TrustedIssuersRegistry{Address: addr, client: client}, nil
}
