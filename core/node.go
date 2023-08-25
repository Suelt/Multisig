package core

import (
	"crypto/ed25519"
)

type Node struct {
	// Fields from Contract

	Mypriv ed25519.PrivateKey
	Mypub  ed25519.PublicKey

	Contractmap map[string]Contract

	Configmap map[string]Config
}

type Contract interface {
	invoke(param []byte, functionname string) string
}

type Config interface {
	generateContract() Contract
}

func NewNode() (*Node, error) {
	pubkey, privkey, err := ed25519.GenerateKey(nil)
	conmap := make(map[string]Contract, 0)
	// con := NewMultiSigContract(string(pubkey))
	return &Node{
		// Con:    con,
		Mypriv:      privkey,
		Mypub:       pubkey,
		Contractmap: conmap,
	}, err
}

// func (n *Node) SignMessage(msg []byte) {

// 	n.Con.SignMessage(msg, n.Mypub)

// 	ifcommitted := n.Con.CheckIfCommitted(msg)
// 	if ifcommitted {
// 		n.Con.Commit(msg)
// 	} else {
// 		fmt.Printf("msg %s has gained %d votes", string(msg), len(n.Con.SignedMap[string(msg)]))
// 	}

// }

// func (n *Node) Publish(msg []byte) {
// 	n.Con.ReceiveMessage(msg)

// }

func (n *Node) Invoke(param []byte, function string, contractaddr string) string {
	c := n.Contractmap[contractaddr]
	return c.invoke(param, function)
}

func (n *Node) Testinvoke() {
	n.Invoke([]byte{1}, "receivemessage", "1")
	n.Invoke([]byte{2}, "signmessage", "1")
}

func (n *Node) NewSmartContract(conf Config) Contract {
	return conf.generateContract()
}
