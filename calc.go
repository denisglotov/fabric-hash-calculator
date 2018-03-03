package main

import (
	"encoding/hex"
	"fmt"
	cb "github.com/hyperledger/fabric/protos/common"
	"log"
	"strings"
)

func strToHex(str string) []byte {
	str = strings.TrimPrefix(str, "0x")
	strBytes := []byte(str)
	dst := make([]byte, hex.DecodedLen(len(strBytes)))
	n, err := hex.Decode(dst, strBytes)
	if err != nil {
		log.Fatal(err)
	}
	return dst[:n]
}

func main() {
	fmt.Print("Enter previous hash: ")
	var prevHashStr string
	fmt.Scanln(&prevHashStr)
	prevHash := strToHex(prevHashStr)

	fmt.Print("Enter data hash: ")
	var dataHashStr string
	fmt.Scanln(&dataHashStr)
	dataHash := strToHex(dataHashStr)

	h := cb.BlockHeader{
		Number:       1,
		PreviousHash: prevHash,
		DataHash:     dataHash}

	fmt.Println("Current block hash: ", hex.EncodeToString(h.Hash()))
}
