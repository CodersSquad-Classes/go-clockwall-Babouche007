package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func readTime(port string) string {

	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error connecting to server:", err)
		os.Exit(1)
	}

	var time string
	_, err = fmt.Fscanln(conn, &time)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading time from server:", err)
		os.Exit(1)
	}

	conn.Close()
	return time
}

func main() {
	length := len(os.Args)
	i := 1
	for length > 1 {
		arg := os.Args[i]

		args := strings.Split(arg, "=")
		if len(args) != 2 || args[0] == "" {
			fmt.Println("Usage: clockWall [Timezone]=localhost:<port>")
			os.Exit(1)
		}

		port := strings.Split(args[1], ":")
		if port[0] != "localhost" {
			fmt.Println("Usage: clockWall [Timezone]=localhost:<port>")
			os.Exit(1)
		}

		if _, err := strconv.Atoi(port[1]); err != nil {
			fmt.Println("Error with port : port should be a number")
			os.Exit(1)
		}

		fmt.Println(args[0] + " : " + readTime(args[1]))

		length--
		i++
	}
}
