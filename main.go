package main

import (
	"fmt"
	"encoding/hex"
	"time"
	"runtime/pprof"
	"os"
	"encoding/binary"
	"runtime"
)

func main() {
	f, _ := os.Create("profile.pp")
	runtime.SetCPUProfileRate(100)
	pprof.StartCPUProfile(f)
	fmt.Println("testing skein")
	sk := new(Skein1024)
	sk.Init(1024)
	out := make([]byte, 128)
	t := time.Now()
	bf := make([]byte, 4)
	for i := 0; i < 1000000; i++ {
		sk.Init(1024)
		binary.LittleEndian.PutUint32(bf, uint32(i))
		sk.Update(bf)
		sk.Final(out)
	}
	pprof.StopCPUProfile()
	fmt.Println(float64(time.Now().UnixNano() - t.UnixNano()) / 1000000000.0)
	fmt.Println(hex.EncodeToString(out))

}

