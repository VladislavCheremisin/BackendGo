package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server has started...")

	for {
		conn, err := listener.Accept()
		fmt.Println("Accept has started")
		if err != nil {
			log.Print(err)
			continue
		}
		go broadcaster(conn)
		who := conn.RemoteAddr().String()
		terminalMas := "Client connected " + who
		fmt.Println(terminalMas)

		go handleConn(conn)
	}
}
func broadcaster(conn net.Conn) {
	clients := make(map[client]bool)
	//clientsName := make(map[int]string)

	type person struct {
		name string
		ip   string
	}

	//var nomberCli int

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			fmt.Println("cli := <-leaving:...")
			delete(clients, cli)
			close(cli)
		}
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Your address is " + who
	messages <- who + " has arrived"
	entering <- ch

	go clientReader(conn)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		messages <- "SERVER: " + input.Text()
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientReader(conn net.Conn) {
	inputCli := bufio.NewScanner(conn)
	who := conn.RemoteAddr().String()
	for inputCli.Scan() {
		messages <- who + " " + inputCli.Text()
	}
}
