package main

import (
	"net"
	"sync"
)

const MaxPacketSize = 4096

var bufPool = sync.Pool {
	New: func() interface{} {
		return make([]byte, MaxPacketSize)
	},
}

func process(outChan chan []byte) {
	for data := range outChan {
		// process data

		

		// Re-slice to maximum capacity and return it
		// for re-use. This is important to guarantee that
		// all calls to Get() will return a buffer of
		// length MaxPacketSize.
		bufPool.Put(data[:MaxPacketSize])
	}
}

func reader(conn net.PacketConn, outChan chan []byte) {
	for {
		data := bufPool.Get().([]byte)

		n, _, err := conn.ReadFrom(data)

		if err != nil {
			break
		}

		outChan <- data[:n]
	}

	close(outChan)
}

func main() {
	N := 3
	var wg sync.WaitGroup

	outChan := make(chan []byte, N)

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			process(outChan)
			wg.Done()
		}()
	}

	wg.Add(1)

	conn, err := net.ListenPacket("udp", "localhost:10001")

	if err != nil {
		panic(err.Error())
	}

	go func() {
		reader(conn, outChan)
		wg.Done()
	}()

	wg.Wait()
}
