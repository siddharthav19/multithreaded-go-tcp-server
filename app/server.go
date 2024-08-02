package main

import (
	"fmt"
	"net"
	"os"
)

/*

An HTTP request/response is made up of three parts, each separated by a CRLF (\r\n):

1. Request/Response line.
2. Zero or more headers, each ending with a CRLF.
3. Optional request/response body. and at the end an escape character \x00

*/

func main() {
	l, err := net.Listen("tcp", "localhost:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221", err)
		os.Exit(1)
	}
	for {
		con, err := l.Accept()
		fmt.Print(con.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
	}

}
