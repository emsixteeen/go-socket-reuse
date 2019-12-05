package main

import (
	"fmt"
	"bufio"
	"github.com/dspinhirne/netaddr-go"
	"github.com/libp2p/go-reuseport"
)

const (
	REMOTE = "%s:9090"
	LOCAL = "127.0.0.1:9091"
)

func connect(conn int, laddr, raddr string, next chan bool) {
	c, err := reuseport.Dial("tcp4", laddr, raddr)
	if err != nil {
		panic(fmt.Sprintf("%T: %s", err, err))
	}
	defer c.Close()
	fmt.Printf("Connected #%d, from %s -> %s\n", conn, laddr, raddr)
	next <- true

	for {
		d, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			panic(fmt.Sprintf("%T: %s", err, err))
		}

		fmt.Printf("Data (Connection #%d) %s\n", conn, d)
	}
}

func main() {
	const MAX = 1024*32
	next := make(chan bool)
	conn := 0
	local := LOCAL
	ip, _ := netaddr.ParseIPv4("127.0.0.1")

	for i := 0; i< MAX; i++ {
		conn++
		remote := fmt.Sprintf(REMOTE, ip)

		go connect(conn, local, remote, next)

		<-next
		ip = ip.Next()
	}

	fmt.Println("Reached the end of the line!")
	<-next
}
