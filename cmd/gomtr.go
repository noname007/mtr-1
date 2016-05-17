package main

import (
	"fmt"
	"github.com/liuxinglanyue/mtr"
)

func main() {
	fmt.Println("hello")
	fmt.Println(mtr.DEFAULT_RETRIES)
	mtr.LocalAddr()
	// host := "gs.push.rgbvr.com"
	destAddrs, _ := mtr.DestAddr("gs.push.rgbvr.com")

	for _, destAddr := range destAddrs {
		fmt.Println(destAddr)
	}

	//
	c := make(chan mtr.TracerouteHop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			fmt.Println(hop.TTL, hop.Address, hop.AvgTime, hop.BestTime, hop.Loss)
		}
	}()
	options := mtr.TracerouteOptions{}
	_, err := mtr.Mtr(destAddrs, &options, c)
	if err != nil {
		fmt.Println(err)
	}
	//

	mm, err := mtr.T("gs.push.rgbvr.com", true, 0, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mm)

	info, err := mtr.T("gs.push.rgbvr.com", false, 0, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)
}
