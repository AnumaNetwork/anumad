package externalapi

import (
	"math/big"
	"reflect"
	"testing"
)

func initTestBlockInfoStructsForClone() []*BlockInfo {

	tests := []*BlockInfo{
		{
			true,
			BlockStatus(0x01),
			0,
			big.NewInt(0),
			nil,
			[]*DomainHash{},
			[]*DomainHash{},
		}, {
			true,
			BlockStatus(0x02),
			0,
			big.NewInt(0),
			nil,
			[]*DomainHash{},
			[]*DomainHash{},
		}, {
			true,
			1,
			1,
			big.NewInt(0),
			nil,
			[]*DomainHash{},
			[]*DomainHash{},
		}, {
			true,
			255,
			2,
			big.NewInt(0),
			nil,
			[]*DomainHash{},
			[]*DomainHash{},
		}, {
			true,
			0,
			3,
			big.NewInt(0),
			nil,
			[]*DomainHash{},
			[]*DomainHash{},
		}, {
			true,
			BlockStatus(0x01),
			0,
			big.NewInt(1),
			nil,
			[]*DomainHash{},
			[]*DomainHash{},
		}, {
			false,
			BlockStatus(0x01),
			0,
			big.NewInt(1),
			NewDomainHashFromByteArray(&[DomainHashSize]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}),
			[]*DomainHash{
				NewDomainHashFromByteArray(&[DomainHashSize]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02}),
				NewDomainHashFromByteArray(&[DomainHashSize]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03}),
			},
			[]*DomainHash{
				NewDomainHashFromByteArray(&[DomainHashSize]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04}),
				NewDomainHashFromByteArray(&[DomainHashSize]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05}),
			},
		},
	}
	return tests
}

func TestBlockInfo_Clone(t *testing.T) {

	blockInfos := initTestBlockInfoStructsForClone()
	for i, blockInfo := range blockInfos {
		blockInfoClone := blockInfo.Clone()
		if !reflect.DeepEqual(blockInfo, blockInfoClone) {
			t.Fatalf("Test #%d:[DeepEqual] clone should be equal to the original", i)
		}
	}
}
