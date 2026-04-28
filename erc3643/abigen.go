// Package erc3643 provides Go bindings + helpers for ERC-3643 contracts.
//
// Source of truth: luxfi/erc-3643 (fork of ERC-3643/ERC-3643 v4.1.3).
package erc3643

//go:generate sh -c "rm -rf ./contracts-vendored && git clone --depth 1 --branch v4.1.3 https://github.com/luxfi/erc-3643 ./contracts-vendored"
//go:generate sh -c "find ./contracts-vendored -name '*.json' -path '*/artifacts/*' | head -20"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/token/Token.sol/Token.json --pkg erc3643 --type Token --out token_bindings.go"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/compliance/modular/ModularCompliance.sol/ModularCompliance.json --pkg erc3643 --type Compliance --out compliance_bindings.go"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/registry/implementation/IdentityRegistry.sol/IdentityRegistry.json --pkg erc3643 --type IdentityRegistry --out identityregistry_bindings.go"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/registry/implementation/TrustedIssuersRegistry.sol/TrustedIssuersRegistry.json --pkg erc3643 --type TrustedIssuersRegistry --out trustedissuersregistry_bindings.go"
//go:generate sh -c "abigen --abi ./contracts-vendored/artifacts/contracts/registry/implementation/ClaimTopicsRegistry.sol/ClaimTopicsRegistry.json --pkg erc3643 --type ClaimTopicsRegistry --out claimtopicsregistry_bindings.go"
