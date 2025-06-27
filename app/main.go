package main

import (
	"fmt"
	"ftp/transfer"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [receive PORT | send FILE IP:PORT]")
		return
	}

	switch os.Args[1] {
	case "recieve":
		if len(os.Args) != 3 {
			fmt.Println("Usage: recieve PORT")
			return
		}
		port := os.Args[2]
		transfer.StartReciever(port)

	case "send":
		if len(os.Args) != 4 {
			fmt.Println("Usage: send FILE IP:PORT")
			return
		}
		file := os.Args[2]
		addr := os.Args[3]
		err := transfer.SendFile(file, addr)
		if err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Unknown command")
	}
}
