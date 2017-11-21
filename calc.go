package main

import (
	cb "github.com/hyperledger/fabric/protos/common"
	"encoding/hex"
	"fmt"
	"log"
)

func strToHex(str string) []byte {
	str_bytes := []byte(str)
	dst := make([]byte, hex.DecodedLen(len(str_bytes)))
	n, err := hex.Decode(dst, str_bytes)
	if err != nil {
		log.Fatal(err)
	}
	return dst[:n]
}

func main() {
	h := cb.BlockHeader{
		Number: 1,
		PreviousHash: strToHex(
			"369ae6cee7c2c2b781cad7d41e7299772b1fb4102c0617d35b4375db4d791a08"),
		DataHash: strToHex(
			"fd27225b06b21846335a5feafba57c1d55eee456bad19723b326f8810defd7d6")}
	fmt.Println("Current block hash: ", hex.EncodeToString(h.Hash()))
}
