package domain

import (
	"errors"
	"strconv"
	"strings"
)

var byte uint = 1
var kib uint = byte * 1024
var mib uint = kib * 1024
var gib uint = mib * 1024
var tib uint = gib * 1024
var pib uint = tib * 1024
var eib uint = pib * 1024

var binaryUnits = map[string]uint{
	"Bytes": byte,
	"KiB":   kib,
	"MiB":   mib,
	"GiB":   gib,
	"TiB":   tib,
	"PiB":   pib,
	"EiB":   eib,
}

func ConvertToBytes(s string) (float64, error) {
	str := strings.Split(s, " ")

	if len(str) != 2 {
		return 0, errors.New("invalid size")
	}

	sizeStr := str[0]
	unitStr := str[1]

	unit := binaryUnits[unitStr]
	floatSize, err := strconv.ParseFloat(sizeStr, 64)

	if err != nil {
		return 0, err
	}

	return floatSize * float64(unit), nil
}
