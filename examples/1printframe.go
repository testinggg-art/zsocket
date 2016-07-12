package main

import (
	"fmt"

	"github.com/nathanjsweet/zsocket"
	"github.com/nathanjsweet/zsocket/nettypes"
)

func main() {
	// args: interface index, options, ring block count, frameOrder, framesInBlock packet types
	// unless you know what you're doing just pay attention to the interface index, whether
	// or not you want the tx ring, rx ring, or both enabled, and what nettype you are listening
	// for.
	zs, err := zsocket.NewZSocket(14, zsocket.ENABLE_RX, 256, zsocket.MAX_ORDER, 4, nettypes.All)
	if err != nil {
		panic(err)
	}
	zs.Listen(func(f *nettypes.Frame, frameLen uint16) {
		fmt.Printf(f.String(frameLen, 0))
	})
}
