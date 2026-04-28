package onchainid

import (
	"context"
	"errors"

	"github.com/luxfi/geth/common"
)

// Client is the minimal ETH-client surface this package needs.
// We keep it interface-shaped so consumers can pass *ethclient.Client
// or any compatible test double.
type Client interface {
	ChainID(ctx context.Context) (any, error)
}

// Identity is an ergonomic wrapper over the abigen-generated Identity contract.
//
// Once `go generate ./...` produces identity_bindings.go, the Bindings field
// will hold the typed contract handle. Until then, callers should treat
// New as a forward-compatible constructor that pins the address + client.
type Identity struct {
	Address common.Address
	client  Client
	// Bindings *IdentityBindings // populated by abigen-generated code
}

// NewIdentity binds to an existing Identity contract at addr.
func NewIdentity(client Client, addr common.Address) (*Identity, error) {
	if client == nil {
		return nil, errors.New("onchainid: nil client")
	}
	if addr == (common.Address{}) {
		return nil, errors.New("onchainid: zero address")
	}
	return &Identity{Address: addr, client: client}, nil
}
