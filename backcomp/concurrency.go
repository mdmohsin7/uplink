// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

// Package backcomp contains utilities for handling backwards incompatible changes.
package backcomp

import (
	"context"

	"storj.io/uplink"
	"storj.io/uplink/internal/expose"
)

// RequestAccessWithPassphraseAndConcurrency requests satellite for a new access using a passhprase and specific concurrency for the Argon2 key derivation.
func RequestAccessWithPassphraseAndConcurrency(ctx context.Context, config uplink.Config, satelliteNodeURL, apiKey, passphrase string, concurrency uint8) (_ *uplink.Access, err error) {
	fn, ok := expose.RequestAccessWithPassphraseAndConcurrency.(func(ctx context.Context, config uplink.Config, satelliteNodeURL, apiKey, passphrase string, concurrency uint8) (_ *uplink.Access, err error))
	if !ok {
		return nil, uplink.Error.New("invalid type %T", fn)
	}
	return fn(ctx, config, satelliteNodeURL, apiKey, passphrase, concurrency)
}