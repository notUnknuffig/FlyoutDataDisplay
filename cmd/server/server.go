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

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		json, err := state.WriteString(object())
		if err != nil {
			log.Fatal(err)
		} else {
			for {
				_, err = fmt.Fprintf(conn, json)
				if err != nil {
					log.Print(err)
					break
				}
				time.Sleep(20 * time.Millisecond)
			}
		}
	}
}

func object() state.FlightData {
	eng := state.JetEngine{
		Throttle:            0.815123,
		IdleThrottle:        0.3,
		HasAfterburner:      true,
		AfterburnerThrottle: 0.156455,
		AlternatorPower:     3123.23451234,
		HydraulicPower:      312.2561,
		NetThrust:           121350.235132,
		FuelFlow:            10.235123,
	}
	pist := state.PistonEngine{
		Throttle:     0.8,
		IdleThrottle: 0.3,
		RPM:          4000,
		FuelFlow:     1.2,
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
		Altitude:      165.151365514,
		Airspeed:      43.161341234,
		Pitch:         5.136514,
		Roll:          0.16135134,
		G:             1.1365341,
		AGL:           0.13561342,
		Heading:       330.12323512,
		Climb:         12.135123,
		Mass:          10000.135123,
		Mach:          0.21353,
		Alpha:         3.05135343,
		Latitude:      -51.51123523,
		Longitude:     -36.313252323,
		JetEngines:    []state.JetEngine{eng, eng, eng, eng, eng},
		PistonEngines: []state.PistonEngine{pist, pist, pist, pist},
		Missiles:      []state.Missile{mis},
		ActiveMissile: "Aim-9X",
	}
}
