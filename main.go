package main

import "multisig/core"

func main() {
	con := core.NewMultiSigContract(nil)
	n, err := core.NewNode()
	if err != nil {
		panic(err)
	}

	n.Contractmap[con.Addr] = con
	n.Invoke([]byte{1}, "ReceiveMessage223", con.Addr)

	//n.SignMessage([]byte("hello"))
}
