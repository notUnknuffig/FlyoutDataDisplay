package main

import (
	"fmt"
	"log"
	"math/rand"
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

		if err != nil {
			log.Fatal(err)
		} else {
			i := 0
			for {
				json, err := state.WriteString(object(i))
				_, err = fmt.Fprintf(conn, json)
				if err != nil {
					log.Print(err)
					break
				}
				i++
				time.Sleep(20 * time.Millisecond)
			}
		}
	}
}

func object(i int) state.FlightData {
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
	var b = false
	if i > 3000 {
		b = true
	}
	tank := state.FuelTank{
		Priority:    1,
		IsEmpty:     b,
		Fuel:        (3000 - float32(i)),
		FuelPercent: (3000 - float32(i)) / 3000,
		Capacity:    3000,
	}

	var randomVal float32
	if i%3600 == 0 {
		randomVal = 0.5 - rand.Float32()
	}

	return state.FlightData{
		Time:          0,
		Name:          "J-8 MFD Test",
		Altitude:      165.151365514 + float32(i),
		Airspeed:      43.161341234 + float32(i)*1,
		Pitch:         5.136514,
		Roll:          (float32(i%1800) * 0.5),
		G:             1.1365341,
		AGL:           0.13561342 + float32(i),
		Heading:       (float32(i%360) * 1),
		Climb:         12.135123,
		Mass:          10000.135123,
		Mach:          0.21353,
		Alpha:         3.05135343 + (float32(i) * (randomVal / float32(i%3600))),
		Latitude:      35.604,  //+ float32(math.Sin(stateNavigation.DegToRad((float64(i%360)*1)))*0.1),
		Longitude:     51.8061, // + float32(math.Cos(stateNavigation.DegToRad((float64(i%360)*1)))*0.1),
		JetEngines:    []state.JetEngine{eng, eng, eng, eng, eng},
		PistonEngines: []state.PistonEngine{pist, pist, pist, pist},
		FuelTanks:     []state.FuelTank{tank, tank, tank},
		Fuel:          6000 - float32(i),
		FuelCapacity:  6000,
		FuelRatio:     (6000 - float32(i*2)) / 6000,
		TimeToEmpty:   13000,
		Missiles:      []state.Missile{mis},
		ActiveMissile: "Aim-9X",
	}
}
