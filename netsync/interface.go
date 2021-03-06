// Copyright (c) 2017 The btcsuite developers
// Copyright (c) 2018-2019 The Soteria DAG developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/totaloutput/soterd/blockdag"
	"github.com/totaloutput/soterd/chaincfg"
	"github.com/totaloutput/soterd/chaincfg/chainhash"
	"github.com/totaloutput/soterd/mempool"
	"github.com/totaloutput/soterd/peer"
	"github.com/totaloutput/soterd/soterutil"
	"github.com/totaloutput/soterd/wire"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *soterutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockdag.BlockDAG
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	// NOTE(cedric): Commented out to disable checkpoint-related code (JIRA DAG-3)
	// 
	//
	// DisableCheckpoints bool
	MaxPeers int

	FeeEstimator *mempool.FeeEstimator
}
