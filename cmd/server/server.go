package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"example.com/MFDTest/internal/state"
)

func main() {
	const addr = "127.0.0.1:14501"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("listening on", addr)

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	json, err := state.WriteString(object())
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err = fmt.Fprintf(conn, json)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func object() state.FlightData {
	eng := state.JetEngine{
		Throttle:            1.0,
		IdleThrottle:        0.3,
		HasAfterburner:      true,
		AfterburnerThrottle: 1.0,
		AlternatorPower:     30000,
		HydraulicPower:      300,
		NetThrust:           120000,
		FuelFlow:            10.0,
	}
	pist := state.PistonEngine{
		Throttle:     1.0,
		IdleThrottle: 0.3,
		RPM:          4000,
		FuelFlow:     0.2,
		Temperature:  400.0,
		Power:        660000,
	}
	mis := state.Missile{
		Count: 1,
		Type:  "Infrared Missile",
		Name:  "Aim-9X",
	}
	return state.FlightData{
		Time:          0,
		Name:          "J-8 MFD Test",
		Altitude:      165.0,
		Airspeed:      43,
		Pitch:         5.0,
		Roll:          0.0,
		G:             1.2,
		AGL:           0.5,
		Heading:       330,
		Climb:         12.0,
		Mass:          10000,
		Mach:          0.2,
		Alpha:         3.05,
		Latitude:      -51.51,
		Longitude:     -36.3,
		JetEngines:    []state.JetEngine{eng},
		PistonEngines: []state.PistonEngine{pist},
		Missiles:      []state.Missile{mis},
		ActiveMissile: "Aim-9X",
	}
}
