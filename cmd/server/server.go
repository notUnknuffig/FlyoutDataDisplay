package main

import (
	"fmt"
	"log"
	"math"
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
	/* mis1 := state.Missile{
		Count:       2,
		Type:        "Infrared Missile",
		Name:        "Aim-9X",
		MaxDistance: 15000.0,
	}
	mis2 := state.Missile{
		Count:       12,
		Type:        "Radar Missile",
		Name:        "ARAAM-120",
		MaxDistance: 10000.0,
	} */
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

	// var randomVal float32
	// if i%3600 == 0 {
	// 	randomVal = 0.5 - rand.Float32()
	// }

	miss := "Aim-9X"
	if i%1200 > 300 && i%1200 < 600 {
		miss = "ARAAM-120"
	} else if i%1200 > 600 {
		miss = "PL-12"
	}

	misss := []state.Missile{
		{
			Count:       1,
			Type:        "InfraredAllAspect",
			Name:        "Aim-9X",
			MaxDistance: 10000.0,
		},
		{
			Count:       4,
			Type:        "CraftRadar",
			Name:        "ARAAM-120",
			MaxDistance: 10000.0,
		}, {
			Count:       6,
			Type:        "Unguided",
			Name:        "PL-12",
			MaxDistance: 10000.0,
		}, {
			Count:       13,
			Type:        "Radar Missile",
			Name:        "Aim-9X",
			MaxDistance: 10000.0,
		},
	}

	return state.FlightData{
		Time:          0,
		Name:          "J-8 MFD Test",
		Altitude:      165.151365514,
		Airspeed:      43.161341234,
		Pitch:         float32(math.Mod(float64(5.136514)+float64(i)*0.05, 180)) - 90,
		Roll:          0.0,
		G:             1.1365341,
		AGL:           0.13561342,
		Heading:       308.01,
		Climb:         12.135123,
		Mass:          10000.135123,
		Mach:          0.21353,
		Alpha:         3.05135343,
		Beta:          0.05135343,
		Latitude:      -34.50,  //+ float32(math.Sin(stateNavigation.DegToRad((float64(i%360)*1)))*0.1),
		Longitude:     -51.780, // + float32(math.Cos(stateNavigation.DegToRad((float64(i%360)*1)))*0.1),
		JetEngines:    []state.JetEngine{eng, eng, eng, eng, eng},
		PistonEngines: []state.PistonEngine{pist, pist, pist, pist},
		FuelTanks:     []state.FuelTank{tank, tank, tank},
		Fuel:          6000 - float32(i),
		FuelCapacity:  6000,
		FuelRatio:     (6000 - float32(i*2)) / 6000,
		TimeToEmpty:   13000,
		Missiles:      misss,
		ActiveMissile: miss,
		Radar: state.Radar{
			Range:  12540.0,
			Mode:   "Bore",
			SteerX: 0.0,
			SteerY: 0.0,
			SteerZ: 0.0,
			Size:   68.0,
		},
	}
}
