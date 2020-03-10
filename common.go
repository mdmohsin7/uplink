// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package uplink

import (
	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"

	"storj.io/common/errs2"
	"storj.io/common/rpc/rpcstatus"
)

var mon = monkit.Package()

// Error is default error class for uplink.
var Error = errs.Class("uplink")

// ErrRequestsLimitExceeded is returned when project will exceeded requests limit per second.
var ErrRequestsLimitExceeded = errs.Class("requests limit exceeded")

// ErrBandwidthLimitExceeded is returned when project will exceeded bandwidth limit.
var ErrBandwidthLimitExceeded = errs.Class("bandwidth limit exceeded")

func convertKnownErrors(err error) error {
	if errs2.IsRPC(err, rpcstatus.ResourceExhausted) {
		// TODO is a better way to do this?
		reErr := errs.Unwrap(err)
		if reErr.Error() == "Exceeded Usage Limit" {
			return ErrBandwidthLimitExceeded.New("")
		} else if reErr.Error() == "Too Many Requests" {
			return ErrRequestsLimitExceeded.New("")
		}
	}

	return Error.Wrap(err)
}