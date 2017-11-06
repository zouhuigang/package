package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {
	encoded := zcrypto.MorseEncodeITU(`HELLO WORLD`)
	fmt.Printf("morse encode %s\n", encoded)

	de, _ := zcrypto.MorseDecodeITU(encoded)
	fmt.Printf("morse decode %s\n", de)

}
