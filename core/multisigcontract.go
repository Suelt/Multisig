package core

import (
	"crypto/ed25519"
	"fmt"
	"reflect"
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

	//TODO: implement genaddr
	addr := "0x01"
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

func (con *MultiSigContract) invoke(param []byte, functionname string) string {
	conForReflect := &MultiSigContract{}
	contracttype := reflect.TypeOf(conForReflect)
	method, exist := contracttype.MethodByName(functionname)
	if exist {
		fmt.Println("function name " + functionname + "exists")
		params := []reflect.Value{reflect.ValueOf(param)}
		retValues := method.Func.Call(params)
		return con.parseReflectValues(retValues)
	} else {
		fmt.Println("function name %s doesn't exist", functionname)
		return "function name doesn't exist" + functionname
	}

	// switch functionname {
	// case "receivemessage":
	// 	con.ReceiveMessage(param)
	// case "signmessage":
	// 	con.SignMessage(param, nil)
	// case "commitmessage":
	// 	con.Commit(param)
	// default:
	// 	fmt.Printf("function name %s undefined", functionname)

	// }

}

func (con *MultiSigContract) parseReflectValues(in []reflect.Value) string {
	return ""
}
