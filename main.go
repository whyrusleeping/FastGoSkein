package main

import (
	"./skein"
	"fmt"
	"encoding/hex"
)

func main() {
	fmt.Println("testing skein")
	sk := new(skein.Skein1024)
	sk.Init(1024)
	sk.Update([]byte("TEST"))
	out := make([]byte, 128)
	sk.Final(out)
	fmt.Println(hex.EncodeToString(out))
}

