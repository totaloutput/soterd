// Copyright (c) 2013-2016 The btcsuite developers
// Copyright (c) 2018-2019 The Soteria DAG developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"github.com/totaloutput/soterd/chaincfg"
	"github.com/totaloutput/soterd/wire"
)

// activeNetParams is a pointer to the parameters specific to the
// currently active soter network.
var activeNetParams = &mainNetParams

// params is used to group parameters for various networks such as the main
// network and test networks.
type params struct {
	*chaincfg.Params
	rpcPort string
}

// mainNetParams contains parameters specific to the main network
// (wire.MainNet).  NOTE: The RPC port is intentionally different than the
// reference implementation because soterd does not handle wallet requests.  The
// separate wallet process listens on the well-known port and forwards requests
// it does not handle on to soterd.  This approach allows the wallet process
// to emulate the full reference implementation RPC API.
var mainNetParams = params{
	Params:  &chaincfg.MainNetParams,
	rpcPort: "8334",
}

// regressionNetParams contains parameters specific to the regression test
// network (wire.TestNet).  NOTE: The RPC port is intentionally different
// than the reference implementation - see the mainNetParams comment for
// details.
var regressionNetParams = params{
	Params:  &chaincfg.RegressionNetParams,
	rpcPort: "18334",
}

// testNet1Params contains parameters specific to the test network (version 1)
// (wire.TestNet1).  NOTE: The RPC port is intentionally different than the
// reference implementation - see the mainNetParams comment for details.
var testNet1Params = params{
	Params:  &chaincfg.TestNet1Params,
	rpcPort: "5071",
}

// simNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var simNetParams = params{
	Params:  &chaincfg.SimNetParams,
	rpcPort: "18556",
}

// netName returns the name used when referring to a soter network.  At the
// time of writing, soterd currently places blocks for testnet version 1 in the
// data and log directory "testnet", which does not match the Name field of the
// chaincfg parameters.  This function can be used to override this directory
// name as "testnet" when the passed active network matches wire.TestNet1.
func netName(chainParams *params) string {
	switch chainParams.Net {
	case wire.TestNet1:
		return "testnet"
	default:
		return chainParams.Name
	}
}
