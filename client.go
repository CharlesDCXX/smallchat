package main

import (
	"fmt"
	"net"
)

func main() {
	// 建立 TCP 连接
	conn, err := net.Dial("tcp", "localhost:8972")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	// 发送数据
	message := "hello"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Failed to send data:", err)
		return
	}
	fmt.Println("Data sent:", message)

	// 接收返回值
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to receive data:", err)
		return
	}

	response := string(buffer[:n])
	fmt.Println("Response received:", response)
}
