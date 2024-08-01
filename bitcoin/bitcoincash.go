package bitcoin

import (
	"regexp"
)

type Bitcoincash struct{}

func (Bitcoincash) ChainName() string {
	return "bitcoincash"
}

func (Bitcoincash) CoinbaseDigest(coinbase string) (string, error) {
	return DoubleSha256(coinbase)
}

func (Bitcoincash) HeaderDigest(header string) (string, error) {
	return Sha256Digest(header)
}

func (Bitcoincash) ShareMultiplier() float64 {
	return 65536
}

func (Bitcoincash) ValidMainnetAddress(address string) bool {
	// Apparently a base58 decode is the best way to validate.. TODO.
	return regexp.MustCompile("^(D|A|9)[a-km-zA-HJ-NP-Z1-9]{33,34}$").MatchString(address)
}

func (Bitcoincash) ValidTestnetAddress(address string) bool {
	return regexp.MustCompile("^(n|2)[a-km-zA-HJ-NP-Z1-9]{33}$").MatchString(address)
}

func (Bitcoincash) MinimumConfirmations() uint {
	return uint(101)
}
