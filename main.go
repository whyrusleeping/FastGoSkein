package main

import (
	"fmt"
	"encoding/hex"
	"time"
	"runtime/pprof"
	"os"
)

func main() {
	f, _ := os.Create("profile.pp")
	pprof.StartCPUProfile(f)
	fmt.Println("testing skein")
	sk := new(Skein1024)
	sk.Init(1024)
	out := make([]byte, 128)
	t := time.Now()
	for i := 0; i < 1000001; i++ {
		sk.Update([]byte("TEST"))
		sk.Final(out)
	}
	pprof.StopCPUProfile()
	fmt.Println(float64(time.Now().UnixNano() - t.UnixNano()) / 1000000000.0)
	fmt.Println(hex.EncodeToString(out))

}

