package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/zizimoos/gocoin/utils"
)

const (
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey    string = "30770201010420540de39b3d0433b2aca59b02db4aa9ed975c1da85acdfac95cc4414eb9637be3a00a06082a8648ce3d030107a14403420004cd5f73debf1cc919216f2d0001c0bcaea054cb7aa6844811c3978f77a77fd6b7fd59dc83e3d1a17c089224b2d902f43d652a031e0ba0eccbe8e1e43c539f7caa"
	signature     string = "a84dbfd783232c87a323a7826cddda2bdae553d21de48b08fb5c2fd6351e466e43cf3495dc7086d3a8de7608f8411e0f7acb2cead7b20c71ad1b0625d72d7e4d"
)

func Start() {
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
