// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"github.com/kaspanet/go-muhash"
	"github.com/AnumaNetwork/anumad/domain/consensus/model/externalapi"
	"github.com/AnumaNetwork/anumad/domain/consensus/utils/blockheader"
	"github.com/AnumaNetwork/anumad/domain/consensus/utils/subnetworks"
	"github.com/AnumaNetwork/anumad/domain/consensus/utils/transactionhelper"
	"math/big"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, //script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x61, 0x6E, 0x75, 0x6D, 0x61, 0x2d, 0x6D, 0x61, 0x69, 0x6E, 0x6e, 0x65, 0x74, // anuma-mainnet
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x03, 0x1e, 0x5f, 0xd4, 0x89, 0xfb, 0xdd, 0x59, 
	0xee, 0xa9, 0x0a, 0xa7, 0x50, 0xe9, 0xbf, 0xcf, 
	0x07, 0xb4, 0x97, 0x22, 0xa8, 0x1a, 0xad, 0xba, 
	0x28, 0x6b, 0x2d, 0x76, 0xd2, 0x54, 0x52, 0xf4,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xf0, 0x2e, 0xca, 0xf2, 0xc2, 0xb7, 0x64, 0x11, 
	0xcb, 0x0d, 0x72, 0xb9, 0x2e, 0x89, 0x64, 0x2a, 
	0x8e, 0x10, 0x3d, 0x6c, 0xb8, 0x27, 0x27, 0x4f, 
	0x21, 0x51, 0x6d, 0x9d, 0xa6, 0x04, 0xcf, 0x32,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x191B9560F14,
		0x1e7fffff,
		0x4e616f72,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x61, 0x6E, 0x75, 0x6D, 0x61, 0x2d, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x74, // anuma-devnet
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xfe, 0x69, 0x45, 0xd1, 0xb2, 0x01, 0xb9, 0x4c, 
	0xe7, 0x2c, 0xbc, 0x25, 0x59, 0xac, 0xb7, 0xc7, 
	0x85, 0xa3, 0x37, 0xff, 0x6f, 0x96, 0xd7, 0x1c, 
	0xa1, 0xba, 0xcb, 0x78, 0x41, 0x6c, 0x51, 0xf0,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x55, 0xae, 0xf8, 0x4f, 0xb1, 0x41, 0x4e, 0x6c, 
	0xb5, 0x6b, 0x69, 0x45, 0x7b, 0x04, 0x55, 0x7b, 
	0xb1, 0xf2, 0x2a, 0x2f, 0x15, 0x4c, 0x89, 0x86, 
	0x74, 0xb5, 0x82, 0x04, 0x46, 0xb2, 0xfb, 0x81,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x11e9db49828,
		525264379,
		0x48e5e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x61, 0x6E, 0x75, 0x6D, 0x61, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, 0x74, // anuma-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xaf, 0x73, 0xa4, 0x1b, 0x53, 0x8c, 0xf5, 0x80, 
	0xc4, 0xd5, 0xe8, 0x70, 0xbf, 0xfe, 0xe0, 0xc4, 
	0x80, 0x1f, 0x1b, 0x51, 0x5b, 0x02, 0xef, 0x30, 
	0x1e, 0xf8, 0xce, 0x2f, 0x5e, 0xfc, 0xda, 0x80,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x39, 0xe1, 0xa2, 0x45, 0x3e, 0xa6, 0x4c, 0x09, 
	0xaa, 0x5c, 0x47, 0x57, 0x75, 0x69, 0xd5, 0x73, 
	0xcf, 0x59, 0x66, 0xad, 0xc1, 0xd0, 0x88, 0x0e, 
	0x86, 0x96, 0xd8, 0x10, 0xf5, 0xec, 0xa3, 0xd0,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x207fffff,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                         // Varint
	0x00,                                                                         // OP-FALSE
	0x61, 0x6E, 0x75, 0x6D, 0x61, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65, 0x74, // anuma-testnet
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x14, 0x68, 0xa1, 0x68, 0xbb, 0x1b, 0x8a, 0x63, 
	0x83, 0x7a, 0x2a, 0x4c, 0x7d, 0x4a, 0xb2, 0xc4, 
	0x6e, 0xa0, 0x01, 0xbd, 0xc9, 0x91, 0x7f, 0x05, 
	0xf4, 0x54, 0x15, 0xdb, 0xd3, 0xc9, 0xbe, 0x7b,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x89, 0xf8, 0x4a, 0xbb, 0xae, 0x14, 0x74, 0x6d, 
	0x8f, 0x7e, 0x7a, 0x11, 0x98, 0xe1, 0x53, 0x4e, 
	0xb3, 0xa4, 0x7e, 0x64, 0x19, 0x54, 0xbe, 0x8f, 
	0x11, 0x7a, 0x19, 0x80, 0x2e, 0x5e, 0x0d, 0x40,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x18BF8390F5A,
		0x1e7fffff,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
