package main

import (
	"fmt"
	"io"
	"bufio"
	"net"
	"github.com/dspinhirne/netaddr-go"
)

const (
	PORT = "%s:9090"
)

func conn(c net.Conn, conn int) {
	defer c.Close()

	for {
		d, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF (Connection #%d)\n", conn)
				return
			}

			panic(fmt.Sprintf("%T: %s", err, err))
		}

		fmt.Printf("Data (Connection #%d): %s\n", conn, d)
	}
}

func server(addr string, count int, next chan bool) {
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		panic(fmt.Sprintf("%T: %s", err, err))
	}
	defer l.Close()
	next <- true

	fmt.Printf("Listening #%d, on %s\n", count, addr)
	for {
		c, err := l.Accept()
		if err != nil {
			panic(fmt.Sprintf("%T: %s", err, err))
		}

		count++
		fmt.Printf("Connection #%d, on %s\n", count, addr)
		go conn(c, count)
	}
}

func main() {
	const MAX = 1024*32
	ip, _ := netaddr.ParseIPv4("127.0.0.1")
	next := make(chan bool)

	for i := 0; i<MAX; i++ {
		addr := fmt.Sprintf(PORT, ip)
		go server(addr, i+1, next)

		<-next
		ip = ip.Next()
	}

	fmt.Println("Waiting...")
	<-next
}
