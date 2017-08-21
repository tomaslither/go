package main

import "net"
import "fmt"

func server(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	for {
		_, err := ln.Accept()
		if err != nil {
			fmt.Errorf("Error!")
		} else {
			fmt.Println("Accepted on port ", port)
		}
	}
}

func main() {
	go server(":8000")
	go server(":8001")
	server(":8002")
}
