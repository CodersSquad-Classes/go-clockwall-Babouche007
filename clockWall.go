package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

//read the time on a specified port
func readTime(port string) string {
	//connect to the server
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error connecting to server:", err)
		os.Exit(1)
	}
	//read the time from the server
	var time string
	_, err = fmt.Fscanln(conn, &time)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading time from server:", err)
		os.Exit(1)
	}

	//close the connection
	conn.Close()
	return time
}

func main() {
	arg := os.Args[1]
	//split arg each =
	args := strings.Split(arg, "=")
	fmt.Println(args[0] + " : " + readTime(args[1]))
}
