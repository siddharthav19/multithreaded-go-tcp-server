package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
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

// make it actually streaming (this wont scale for huge file sizes)
func streamFile(con net.Conn, file_param string) {
	path := getFilePath(file_param)
	data, err := os.Stat(path)
	if err != nil {
		con.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}
	file_size := data.Size()
	content, _ := os.ReadFile(path)
	con.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", file_size, content)))
}

func createAndWriteFile(con net.Conn, file_param string, body string) {
	path := getFilePath(file_param)
	file, err := os.Create(path)
	if err != nil {
		con.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}
	writer := bufio.NewWriter(file)
	writer.Write([]byte(body))
	writer.Flush()
	defer file.Close()
	con.Write([]byte("HTTP/1.1 201 Created\r\n\r\n"))
}

func getFilePath(file_param string) string {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "tmp", file_param+".txt")
	return path
}

func extractMethod(request_line string) string {
	method := strings.SplitN(request_line, " ", 2)[0]
	return method
}

func extractRequestBody(req_string string) string {
	last_crlf_index := strings.LastIndex(req_string, "\r\n")
	raw_body := req_string[last_crlf_index:max(len(req_string)-1, 0)]
	body := strings.SplitN(raw_body, "\r\n", 2)[1]
	return strings.Trim(body, "\x00")
}

func handleRequest(con net.Conn) {
	req := make([]byte, 1024)
	con.Read(req)
	req_string := string(req)
	request_line := strings.SplitN(req_string, "\r\n", 2)[0]
	url_path := strings.SplitN(request_line, " ", 3)[1]
	http_method := extractMethod(request_line)
	if strings.Contains(url_path, "/echo/") {
		param := strings.SplitAfter(url_path, "/echo/")[1]
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(param), param)
		con.Write([]byte(response))
	} else if strings.Contains(url_path, "/files/") { // can be made into strategy as we have different algorithms for different methods
		param := strings.SplitAfter(url_path, "/files/")[1]
		if strings.Compare(http_method, "GET") == 0 {
			streamFile(con, param)
		} else if strings.Compare(http_method, "POST") == 0 {
			body := extractRequestBody(req_string)
			fmt.Println(body)
			createAndWriteFile(con, param, body)
		}
	} else if strings.Contains(url_path, "/user-agent") {
		user_agent_suffix := strings.SplitN(req_string, "User-Agent: ", 2)[1]
		user_agent := strings.SplitN(user_agent_suffix, "\r\n", 2)[0]
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(user_agent), user_agent)
		con.Write([]byte(response))
	} else if strings.Compare(url_path, "/") == 0 {
		con.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		con.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
	con.Close()
}
