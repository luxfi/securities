package onchainid

import (
	"os"
	"testing"

	"github.com/luxfi/geth/common"
)

type fakeClient struct{}

func (fakeClient) ChainID(_ any) (any, error) { return nil, nil }

func TestNewIdentityValidatesInputs(t *testing.T) {
	if _, err := NewIdentity(nil, common.HexToAddress("0x1")); err == nil {
		t.Fatal("expected error for nil client")
	}
	if _, err := NewIdentity(struct{ Client }{}, common.Address{}); err == nil {
		t.Fatal("expected error for zero address")
	}
}

func TestIntegrationRequiresRPC(t *testing.T) {
	if os.Getenv("LUX_TESTRPC") == "" {
		t.Skip("requires LUX_TESTRPC env")
	}
}
