//聊天室服务端
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 这个程序中有四种goroutine。
// main和broadcaster各自是一个goroutine 实例，
// 每一个客户端的连接都会有一个handleConn和clientWriter的goroutine。
// broadcaster是 select用法的不错的样例，因为它需要处理三种不同类型的消息。
func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

type client chan<- string //客户端
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan client)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				// 修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好写
				if !clients[cli] {
					continue
				}
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWrite(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "you are " + who
	messages <- who + " is comming"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + " : " + input.Text()
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
