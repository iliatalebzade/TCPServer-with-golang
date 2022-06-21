package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", "192.168.31.124:8888")
	if err != nil {
		log.Fatal("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("started server at :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
