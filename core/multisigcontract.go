package core

import (
	"crypto/ed25519"
	"fmt"
)

type MultiSigContract struct {
	// Fields from Contract

	PubkeySet []ed25519.PublicKey
	Quorum    uint32
	Leader    ed25519.PublicKey
	Addr      string

	MessageReceived map[string][]byte
	SignedMap       map[string]map[*ed25519.PublicKey]bool
	CommittedMap    map[string]bool
}

type MultiSigConfig struct {
	PubkeySet []ed25519.PublicKey
	Quorum    uint32
	Leader    ed25519.PublicKey
}

func NewMultiSigContract(conf *MultiSigConfig) *MultiSigContract {
	con := &MultiSigContract{
		PubkeySet: conf.PubkeySet,
		Quorum:    conf.Quorum,
		Leader:    conf.Leader,
	}
	con.MessageReceived = make(map[string][]byte)
	con.SignedMap = make(map[string]map[*ed25519.PublicKey]bool)
	con.CommittedMap = make(map[string]bool)

	con.Addr = con.GenAddr()
	return con
}

func (con *MultiSigContract) GenAddr() string {
	var addr string
	//TODO: implement genaddr

	return addr
}

func (con *MultiSigContract) ReceiveMessage(msg []byte) {
	con.MessageReceived[string(msg)] = msg
}

func (con *MultiSigContract) SignMessage(msg []byte, pubkey ed25519.PublicKey) error {

	if con.SignedMap[string(msg)] == nil {
		con.SignedMap[string(msg)] = make(map[*ed25519.PublicKey]bool)
	}
	cansign := con.HasRegistered(pubkey)
	if cansign {
		con.SignedMap[string(msg)][&pubkey] = true
	} else {
		return fmt.Errorf("pubkey %s not registered", pubkey)
	}
	return nil
}

func (con *MultiSigContract) CheckIfCommitted(msg []byte) bool {
	if len(con.SignedMap[string(msg)]) < int(con.Quorum) {
		return false
	} else {
		return true
	}
}

func (con *MultiSigContract) Commit(msg []byte) {
	con.CommittedMap[string(msg)] = true
}

func (con *MultiSigContract) HasRegistered(pubkey ed25519.PublicKey) bool {
	for _, v := range con.PubkeySet {
		if v.Equal(pubkey) {
			return true
		}
	}
	return false
}
