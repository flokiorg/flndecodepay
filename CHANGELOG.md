# Changelog

## [1.0.5]

### Changed

- Updated every flokiorg dependency to its current release: `flnd` v0.2.3,
  `go-flokicoin` v0.26.3, `walletd` v0.2.2 and `flokicoin-neutrino` v0.17.2.
- Built with Go 1.26.8, up from 1.26.5, which closes four reachable stdlib
  vulnerabilities (GO-2026-6218 `net/url`, GO-2026-6090 `crypto/tls`,
  GO-2026-5972 `encoding/asn1`, GO-2026-5026 `net/http`).

## [1.0.4]

Dependency-only update: brings in `flnd` v0.2.0-beta and `go-flokicoin`
v0.26.0-alpha, with their transitive `walletd` and `flokicoin-neutrino` bumps.
No code changes in this repo.

### Changed

- Picked up [flnd v0.2.0-beta](https://github.com/flokiorg/flnd/releases/tag/v0.2.0-beta)
  and [go-flokicoin v0.26.0-alpha](https://github.com/flokiorg/go-flokicoin/releases/tag/v0.26.0-alpha).

## [1.0.3]

### Dependency Updates

- Updated dependencies to align with the new `flnd v0.1.21-beta` release, which includes Taproot channel support.
- Routine `go mod tidy` cleanup.

## [1.0.2]

### Dependency Updates

- Updated dependencies to align with `flnd v0.1.20-beta` and `go-flokicoin v0.25.13-alpha`.
- Routine `go mod tidy` cleanup.

## [1.0.1]

### Dependency Updates

- Updated transitive dependency on `github.com/flokiorg/flnd` following its `v0.1.19-beta` release, which migrates the onion routing layer to the native `github.com/flokiorg/lightning-onion` package.
- Routine `go mod tidy` cleanup.

## [1.0.0]

Initial release of flndecodepay - a Go library for decoding Flokicoin Lightning invoices.

### Features

- Decode BOLT-11 Lightning invoices for the Flokicoin network
- Parse invoice components: amount, description, expiry, payment hash, etc.
- Support for Flokicoin-specific network prefixes
