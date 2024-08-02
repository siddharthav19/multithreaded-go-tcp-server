package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(con)
	}

}

func handleRequest(con net.Conn) {
	req := make([]byte, 1024)
	con.Read(req)
	req_string := string(req)
	request_line := strings.SplitN(req_string, "\r\n", 2)[0]
	url_path := strings.SplitN(request_line, " ", 3)[1]
	if strings.Contains(url_path, "/echo/") {
		param := strings.SplitAfter(url_path, "/echo/")[1]
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)
		con.Write([]byte(response))
	} else {
		con.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
	con.Close()
}
