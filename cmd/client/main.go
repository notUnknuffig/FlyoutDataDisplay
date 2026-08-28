package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"example.com/MFDTest/internal/app"
	"example.com/MFDTest/internal/state"
)

const ADDR = "127.0.0.1:14501"

func main() {
	app := app.App{}
	go initTCPConnection(app)
	app.Init()
}

func initWindow() {

}

func initTCPConnection(a app.App) {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", ADDR)
		if err != nil {
			time.Sleep(5 * time.Second)
			fmt.Println("Reattempting to connect to socket")
		} else {
			break
		}
	}
	defer conn.Close()
	fmt.Println("Connected to socket")

	ms := 0
	for {
		ms = time.Now().Nanosecond()
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		msg := string(buf[:n])
		obj := state.FlightData{}
		if err := state.ReadString(msg, &obj); err == nil {
			state.GlobalFlightData = &obj
		} else {
			log.Fatal(err)
		}
		ms = time.Now().Nanosecond() - ms
		state.SmoothData(ms)
	}
}
