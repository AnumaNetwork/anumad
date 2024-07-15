package utils

import (
	"fmt"

	"github.com/AnumaNetwork/anumad/domain/consensus/utils/constants"
)

// FormatAnum takes the amount of sompis as uint64, and returns amount of ANUM with 8  decimal places
func FormatAnum(amount uint64) string {
	res := "                   "
	if amount > 0 {
		res = fmt.Sprintf("%19.8f", float64(amount)/constants.SompiPerAnuma)
	}
	return res
}
