package core

import (
	"crypto/ed25519"
	"fmt"
)

type Node struct {
	// Fields from Contract
	Con *MultiSigContract

	Mypriv ed25519.PrivateKey
	Mypub  ed25519.PublicKey
}

func NewNode() (*Node, error) {
	pubkey, privkey, err := ed25519.GenerateKey(nil)

	// con := NewMultiSigContract(string(pubkey))
	return &Node{
		// Con:    con,
		Mypriv: privkey,
		Mypub:  pubkey,
	}, err
}

func (n *Node) SignMessage(msg []byte) {

	n.Con.SignMessage(msg, n.Mypub)

	ifcommitted := n.Con.CheckIfCommitted(msg)
	if ifcommitted {
		n.Con.Commit(msg)
	} else {
		fmt.Printf("msg %s has gained %d votes", string(msg), len(n.Con.SignedMap[string(msg)]))
	}

}

func (n *Node) Publish(msg []byte) {
	n.Con.ReceiveMessage(msg)

}
