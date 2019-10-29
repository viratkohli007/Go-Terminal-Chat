package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	// "reflect"
	"bufio"
	"os"
)

var upgrader websocket.Upgrader

var dialer = websocket.Dialer{}
var header http.Header
var globalConn *websocket.Conn

func main() {
	fmt.Println("Hello websockets222222")
	conn, _, err := dialer.Dial("ws://localhost:8081/ws", header)
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(reflect.TypeOf(conn))
	globalConn = conn
	// upgrader.CheckOrigin = func(r *http.Request) bool {
	// 	return true
	// }
	// http.HandleFunc("/ws", handler)
	// err := http.ListenAndServe(":8082", nil)
	// if err != nil {
	// 	log.Fatal("Error in creating server")
	// }

	// fmt.Println(conn)
	// fmt.Println(htpResp)
	_, msgByte, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("error in reading ", err)
	}
	fmt.Println("messages from server---->", string(msgByte))
	handler()
	// http.HandleFunc("/ws", handler)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("conn errror", err)
	}

}

func handler() {
	// conn, err := upgrader.Upgrade(w, r, nil)

	for {

		var input string
		fmt.Println("type msg: ")
		// fmt.Scanln(&input)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			// fmt.Printf("Input was: %q\n", line)
		}
		err := globalConn.WriteMessage(1, []byte(input))
		if err != nil {
			fmt.Println("err-->", err)
		}
		_, msgByte, err := globalConn.ReadMessage()
		if err != nil {
			fmt.Println("error in reading ", err)
		}
		fmt.Println("messages from server---->", string(msgByte))
	}
}

// ... Use conn to send and receive messages.
// msgType, msgByte, err := conn.ReadMessage()
// fmt.Println(string(msgByte))
// fmt.Println(msgType)
// if err = conn.WriteMessage(msgType, []byte("hello from server")); err != nil {
// 	fmt.Println("err-->", err)
// }
