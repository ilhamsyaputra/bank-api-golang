package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateNoRekening() string {
	rand.Seed(time.Now().UnixNano())
	bankDigit := 99900000000		// 999
	min := 10000000
	max := 99999999

	return strconv.Itoa(bankDigit + rand.Intn(max - min + 1) + min)
}