package erc3643

import (
	"os"
	"testing"

	"github.com/luxfi/geth/common"
)

func TestNewTokenValidatesInputs(t *testing.T) {
	if _, err := NewToken(nil, common.HexToAddress("0x1")); err == nil {
		t.Fatal("expected error for nil client")
	}
	if _, err := NewToken(struct{ Client }{}, common.Address{}); err == nil {
		t.Fatal("expected error for zero address")
	}
}

func TestIntegrationRequiresRPC(t *testing.T) {
	if os.Getenv("LUX_TESTRPC") == "" {
		t.Skip("requires LUX_TESTRPC env")
	}
}
