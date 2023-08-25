package core

import (
	"crypto/ed25519"
)

type MultiSigConfig struct {
	PubkeySet []ed25519.PublicKey
	Quorum    uint32
	Leader    ed25519.PublicKey
}

type OrderBookConfig struct {
}

func (multiconf *MultiSigConfig) generateContract() Contract {
	return NewMultiSigContract(multiconf)
}

func (obookconf *OrderBookConfig) generateContract() Contract {
	return NewOrderBookContract(obookconf)
}
