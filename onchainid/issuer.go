package onchainid

import (
	"errors"

	"github.com/luxfi/geth/common"
)

// ClaimIssuer wraps the ClaimIssuer contract — used by KYC/AML providers to
// sign claims attached to investor Identities.
type ClaimIssuer struct {
	Address common.Address
	client  Client
}

// NewClaimIssuer binds to an existing ClaimIssuer at addr.
func NewClaimIssuer(client Client, addr common.Address) (*ClaimIssuer, error) {
	if client == nil {
		return nil, errors.New("onchainid: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("onchainid: zero address")
	}
	return &ClaimIssuer{Address: addr, client: client}, nil
}
