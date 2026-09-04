package app

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"example.com/MFDTest/internal/stateNavigation"
)

func readAreaData() ([]stateNavigation.MappedObject, []stateNavigation.MappedObject, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("WARNING: Unable read from '%%appdata%%\\..\\LocalLow\\Stonext Games\\Flyout\\AreaData'. Airports and Objects wont be displayed.\n")
		return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
	}
	path, err := filepath.Abs(filepath.Join(dir, "..", "LocalLow", "Stonext Games", "Flyout", "AreaData", "Areas.txt"))
	if err != nil {
		fmt.Printf("WARNING: Unable to reach '%%appdata%%\\..\\LocalLow\\Stonext Games\\Flyout\\AreaData'. Airports and Objects wont be displayed.\n")
		return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
	}

	return parseAreaData(path)
}

func parseAreaData(path string) ([]stateNavigation.MappedObject, []stateNavigation.MappedObject, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("WARNING: Unable to locate Areas.txt from '%%appdata%%\\..\\LocalLow\\Stonext Games\\Flyout\\AreaData'. Airports and Objects wont be displayed.\n")
		return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
	}
	defer f.Close()

	/* areaString, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("WARNING: Unable to locate Areas.txt from '%%appdata%%\\..\\LocalLow\\Stonext Games\\Flyout\\AreaData'. Airports and Objects wont be displayed.\n")
		return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
	} */

	// fmt.Println("------------------------- Area Data -------------------------")
	// fmt.Println(string(areaString))
	// fmt.Println("------------------------- Parsing Area Data -------------------------")

	var Airfields []stateNavigation.MappedObject
	var Objects []stateNavigation.MappedObject
	scanner := bufio.NewScanner(f)

	depth := 0
	header := ""

	// Area Data
	airfield := false
	currentObj := stateNavigation.MappedObject{
		Type:   stateNavigation.OBJECT,
		Allied: true,
	}

	// Area Object Data
	isStartLoc := false
	pos := []float32{0, 0, 0}
	rot := 0.0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch line {
		case "":
			continue
		case "{":
			depth++
			continue
		case "}":
			if depth >= 2 && header == "Obj" {
				if isStartLoc {
					currentObj.Heading = float32(stateNavigation.RadToDeg(rot))
				}
				isStartLoc = false
				pos = []float32{0, 0, 0}
				rot = 0.0
			} else if depth == 1 {
				if airfield {
					currentObj.Type = stateNavigation.AIRFIELD
					Airfields = append(Airfields, currentObj)
				} else {
					Objects = append(Objects, currentObj)
				}
				airfield = false
				currentObj = stateNavigation.MappedObject{
					Type:   stateNavigation.OBJECT,
					Allied: true,
				}
			}
			depth--
			continue
		}

		// Skip for any material properties.
		if depth >= 3 {
			continue
		}

		if !strings.Contains(line, "=") {
			header = line
			continue
		} else {
			parts := strings.SplitN(line, "=", 2)
			key := strings.ToLower(strings.TrimSpace(parts[0]))
			val := strings.TrimSpace(parts[1])

			if depth == 1 && header == "Area" {
				switch key {
				case "alt":
					floatVal, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println("Unable to read number in 'Areas.txt', aborting.")
						return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
					}
					currentObj.Altitude = float32(floatVal)
				case "lat":
					floatVal, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println("Unable to read number in 'Areas.txt', aborting.")
						return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
					}
					currentObj.Latitude = float32(floatVal) - 90
				case "lon":
					floatVal, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println("Unable to read number in 'Areas.txt', aborting.")
						return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
					}
					currentObj.Longitude = float32(floatVal) - 90
				case "name":
					currentObj.Name = val
				case "start":
					if val == "true" {
						airfield = true
					}
				}
			} else if depth == 2 && header == "Obj" {
				switch key {
				case "name":
					if val == "StartPoint" {
						isStartLoc = true
					}
				case "pos":
					str := strings.Split(val, ",")
					for i := 0; i < len(str); i++ {
						floatVal, err := strconv.ParseFloat(str[i], 64)
						if err != nil {
							fmt.Println("Unable to read number in 'Areas.txt', aborting.")
							return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
						}
						pos[i] = float32(floatVal)
					}
				case "rot":
					str := strings.Split(val, ",")
					if len(str) < 2 {
						fmt.Println("Unable to read number in 'Areas.txt', aborting.")
						return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
					}
					floatVal, err := strconv.ParseFloat(str[1], 64)
					if err != nil {
						fmt.Println("Unable to read number in 'Areas.txt', aborting.")
						return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
					}
					rot = floatVal
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return []stateNavigation.MappedObject{}, []stateNavigation.MappedObject{}, err
	}
	return Airfields, Objects, nil

}
