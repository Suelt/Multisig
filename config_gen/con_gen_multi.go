package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viperRead := viper.New()

	viperRead.AddConfigPath("./")
	viperRead.SetConfigName("conf_gen")

	err := viperRead.ReadInConfig()

	if err != nil {
		fmt.Println(err)

	}

	pubkeynum := viperRead.GetInt("pubkeynum")
	pubkeymap := make(map[string]string, pubkeynum)

	for i := 0; i < pubkeynum; i++ {
		pubkey, _, err := ed25519.GenerateKey(nil)
		if err != nil {
			fmt.Println(err)
		}
		pubkeymap["node"+string(i)] = hex.EncodeToString(pubkey)
	}

}
