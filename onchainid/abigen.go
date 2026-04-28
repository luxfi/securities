// Package onchainid provides Go bindings + helpers for ONCHAINID contracts.
//
// Source of truth: luxfi/onchain-id (fork of onchain-id/solidity v2.2.2-beta3).
package onchainid

//go:generate sh -c "rm -rf ./contracts-vendored && git clone --depth 1 --branch v2.2.2-beta3 https://github.com/luxfi/onchain-id ./contracts-vendored"
//go:generate sh -c "find ./contracts-vendored -name '*.json' -path '*/artifacts/*' | head -20"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/Identity.sol/Identity.json --pkg onchainid --type Identity --out identity_bindings.go"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/factory/IdFactory.sol/IdFactory.json --pkg onchainid --type IdFactory --out idfactory_bindings.go"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/ClaimIssuer.sol/ClaimIssuer.json --pkg onchainid --type ClaimIssuer --out claimissuer_bindings.go"
