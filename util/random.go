package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	curr := []string{"EUR", "USD", "HUF", "GBP", "CHF", "CAD"}
	return curr[rand.Intn(len(curr))]
}

func RandomOwner(owners [][]string) string {
	i := rand.Intn(len(owners))
	owner := owners[i][0]
	return owner
}
