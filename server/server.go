package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":9000") // Listen on port 9000
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("TCP server started on port 9000")

	conn, err := ln.Accept() // Wait for incoming connections
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Client connected from:", conn.RemoteAddr())

	// Create  creates a new file to save the received data
	file, err := os.Create("receive.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Copy the received data into the file
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Error copying data to file:", err)
		return
	}

	fmt.Println("File received and saved successfully")
}
