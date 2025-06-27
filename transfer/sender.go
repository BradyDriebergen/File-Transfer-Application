package transfer

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(filePath, addr string) error {
	// Set up a new dial connection
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("connect error: %w", err)
	}
	defer conn.Close()

	// Open desired file to send
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("file open error: %w", err)
	}
	defer file.Close()

	stat, _ := file.Stat()
	fileName := stat.Name()
	fileSize := stat.Size()

	writer := bufio.NewWriter(conn)

	// Send file name length, file name, and the size of the file
	binary.Write(writer, binary.BigEndian, uint16(len(fileName)))
	writer.Write([]byte(fileName))
	binary.Write(writer, binary.BigEndian, fileSize)

	_, err = io.Copy(writer, file)
	if err != nil {
		return fmt.Errorf("send error: %w", err)
	}

	writer.Flush()
	fmt.Printf("Sent file: %s (%d bytes)\n", fileName, fileSize)
	return nil
}
