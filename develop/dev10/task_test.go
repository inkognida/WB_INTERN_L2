package main

// server for testing

//import (
//	"bufio"
//	"fmt"
//	"io"
//	"log"
//	"net"
//)
//
//func main() {
//	ln, err := net.Listen("tcp", "localhost:8080")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	conn, _ := ln.Accept()
//
//	for {
//		mes, err := bufio.NewReader(conn).ReadString('\n')
//		if err == io.EOF {
//			return
//		}
//		fmt.Print("serv:", mes)
//		_, _ = conn.Write([]byte(mes + "\n"))
//	}
//}