package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// AnumaMainnetPrivate is the version that is used for
// anuma mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var AnumaMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// AnumaMainnetPublic is the version that is used for
// anuma mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var AnumaMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// AnumaTestnetPrivate is the version that is used for
// anuma testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var AnumaTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// AnumaTestnetPublic is the version that is used for
// anuma testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var AnumaTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// AnumaDevnetPrivate is the version that is used for
// anuma devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var AnumaDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// AnumaDevnetPublic is the version that is used for
// anuma devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var AnumaDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// AnumaSimnetPrivate is the version that is used for
// anuma simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var AnumaSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// AnumaSimnetPublic is the version that is used for
// anuma simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var AnumaSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case AnumaMainnetPrivate:
		return AnumaMainnetPublic, nil
	case AnumaTestnetPrivate:
		return AnumaTestnetPublic, nil
	case AnumaDevnetPrivate:
		return AnumaDevnetPublic, nil
	case AnumaSimnetPrivate:
		return AnumaSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case AnumaMainnetPrivate:
		return true
	case AnumaTestnetPrivate:
		return true
	case AnumaDevnetPrivate:
		return true
	case AnumaSimnetPrivate:
		return true
	}

	return false
}
