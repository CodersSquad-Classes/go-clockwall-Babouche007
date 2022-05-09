// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if os.Args[1] != "-port" {
		fmt.Println("Usage: clockServer -port <port>")
		os.Exit(1)
	}
	if _, err := strconv.Atoi(os.Args[2]); err != nil {
		fmt.Println("Error with port : port should be a number")
		os.Exit(1)
	}
	port := "localhost:" + os.Args[2]
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}

}
