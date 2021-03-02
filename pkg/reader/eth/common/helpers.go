package common

import (
	"log"
	"strconv"
	"strings"
)

func RemoveHexPrefix(hex string) string {
	hex = strings.ToLower(hex)
	if strings.HasPrefix(hex, "0x") {
		hex = strings.TrimLeft(hex, "0x")
	}

	return hex
}

func WeiToEth(wei int64) float64 {
	return float64(wei) / float64(1000000000000000000)
}

func DecodeHex(hex string) int64 {
	formatted, err := strconv.ParseInt(RemoveHexPrefix(hex), 16, 64)
	if nil != err {
		log.Fatal("ETH cant be converted ERROR:", err)
	}
	return formatted
}
