package transfer

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func StartReciever(port string) {
	// Starts listener for a reciever
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	fmt.Println("Server listening on port: ", port)

	for {
		// Checks for connection
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(connection)
	}
}

/*
Once a connection is made, scan metadata.
*/
func handleConnection(conn net.Conn) {
	// Waits to close connection until the function returns
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Gets the name nength from metadata
	var nameLen uint16
	binary.Read(reader, binary.BigEndian, &nameLen)

	// Makes a byte array, reads the filename, saves it as string
	fileNameBytes := make([]byte, nameLen)
	reader.Read(fileNameBytes)
	fileName := string(fileNameBytes)

	// Gets the file size from metadata
	var fileSize int64
	binary.Read(reader, binary.BigEndian, &fileSize)

	// Create file to store new file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Write error: ", err)
	}
	defer file.Close()

	// Copies the byte array from reader and writes it to file
	written, err := io.CopyN(file, reader, fileSize)
	if err != nil {
		fmt.Println("Write error: ", err)
	}
	fmt.Printf("Recieved file: %s (%d bytes)\n", fileName, written)
}
