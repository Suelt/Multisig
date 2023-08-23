package main

import (
	"cyd/core"
)

func main() {
	con := core.NewMultiSigContract(nil)
	n, err := core.NewNode()
	if err != nil {
		panic(err)
	}
	n.Con = con
	n.SignMessage([]byte("hello"))
}
