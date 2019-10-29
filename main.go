package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var upgrader websocket.Upgrader
var dialer = websocket.Dialer{}
var header http.Header

func main() {
	fmt.Println("Hello websockets")

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	http.HandleFunc("/ws", handler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error in creating server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//    conn2, _, err := dialer.Dial("ws://localhost:8082/ws", header)
	// if err != nil {
	// 		    log.Println(err)
	// 		    return
	// 		}
	for {
		// fmt.Scanf("Write your message: ")
		var input string
		fmt.Println("Enter server message:")
		// fmt.Scanln(&input)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			// fmt.Printf("Input was: %q\n", line)
		}
		err = conn.WriteMessage(1, []byte(input))
		if err != nil {
			fmt.Println("err-->", err)
		} else {

			// fmt.Println(conn)
			// fmt.Println(htpResp)

			_, msgByte, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("error in reading ", err)
			}
			fmt.Println("messages from client---->", string(msgByte))
			// if err := conn2.Close(); err != nil {
			// 	fmt.Println("error in closing")
			// }

		}
	}
	// ... Use conn to send and receive messages.
	msgType, msgByte, err := conn.ReadMessage()
	fmt.Println(string(msgByte))
	fmt.Println(msgType)
	if err = conn.WriteMessage(1, []byte("hello from server")); err != nil {
		fmt.Println("err-->", err)
	}
}
