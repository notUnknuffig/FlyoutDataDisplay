package navPointEntry

import (
	"math"
	"strconv"
	"strings"

	"example.com/MFDTest/internal/navigation"
	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
	navMan        *navigation.NavManager
	returnState   state.State
	adding        bool
	navPointIndex int
	navPoint      navigation.MappedObject

	valueIndex  int
	numberIndex int

	lat          []int
	latNegative  bool
	long         []int
	longNegative bool
	height       []int
}

func Init(navMan *navigation.NavManager, returnState state.State) _state {
	s := _state{
		navMan:      navMan,
		returnState: returnState,

		valueIndex:  0,
		numberIndex: 0,
	}
	return s.enterEdit()
}

func (s _state) enterEdit() _state {
	if s.navMan.SelectedObject == -1 {
		return s.enterAdd()
	}
	s.navPointIndex = s.navMan.SelectedObject
	s.navPoint = s.navMan.NavPoints[s.navMan.SelectedObject]
	s.latNegative = s.navPoint.Latitude < 0
	s.lat = splitNumber(s.navPoint.Latitude)
	s.longNegative = s.navPoint.Longitude < 0
	s.long = splitNumber(s.navPoint.Longitude)
	s.height = splitNumber(s.navPoint.Altitude * float32(state.GlobalOptions.Units.HightConversion))
	return s
}

func (s _state) enterAdd() _state {
	s.navPoint = navigation.MappedObject{
		Name:      "Nav Point " + strconv.FormatInt(int64(s.navPointIndex+1), 10),
		Latitude:  state.GlobalFlightData.Latitude,
		Longitude: state.GlobalFlightData.Longitude,
		Heading:   state.GlobalFlightData.Heading,
		Altitude:  state.GlobalFlightData.Altitude,
		Allied:    true,
		Type:      navigation.NAV_POINT,
	}
	s.navPointIndex = s.navMan.SelectedObject
	s.latNegative = s.navPoint.Latitude < 0
	s.lat = splitNumber(s.navPoint.Latitude)
	s.longNegative = s.navPoint.Longitude < 0
	s.long = splitNumber(s.navPoint.Longitude)
	s.height = splitNumber(s.navPoint.Altitude * float32(state.GlobalOptions.Units.HightConversion))
	return s
}

func (s _state) Input() state.State {
	if rl.IsKeyPressed(state.KEY_F8) {
		s.navPoint = s.readNavPoint()
		if s.navPointIndex < 0 {
			s.addNavPoint()
		} else {
			s.navMan.NavPoints[s.navPointIndex] = s.navPoint
		}
		return s.returnState
	} else if rl.IsKeyPressed(state.KEY_F9) {
		s.navPoint = s.readNavPoint()
		s.addNavPoint()
		return s.returnState
	} else if rl.IsKeyPressed(state.KEY_F10) {
		s.removeNavPoint()
		return s.returnState
	}

	if rl.IsKeyPressed(state.KEY_F11) && s.numberIndex > 0 {
		s.numberIndex--
	} else if rl.IsKeyPressed(state.KEY_F11) && s.numberIndex == 0 && s.valueIndex > 0 {
		s.numberIndex = 6
		s.valueIndex--
	} else if rl.IsKeyPressed(state.KEY_F12) && s.numberIndex < 6 {
		s.numberIndex++
	} else if rl.IsKeyPressed(state.KEY_F12) && s.numberIndex == 6 && s.valueIndex < 2 {
		s.numberIndex = 0
		s.valueIndex++
	}

	if rl.IsKeyPressed(state.KEY_F6) || rl.IsKeyPressed(state.KEY_F7) {
		if s.numberIndex == 0 {
			switch s.valueIndex {
			case 0:
				s.latNegative = !s.latNegative
			case 1:
				s.longNegative = !s.longNegative
			case 2:
				if rl.IsKeyPressed(state.KEY_F6) && s.height[s.numberIndex] < 9 {
					s.height[s.numberIndex]++
				} else if rl.IsKeyPressed(state.KEY_F7) && s.height[s.numberIndex] > 0 {
					s.height[s.numberIndex]--
				} else if rl.IsKeyPressed(state.KEY_F6) && s.height[s.numberIndex] == 9 {
					s.height[s.numberIndex] = 0
				} else if rl.IsKeyPressed(state.KEY_F7) && s.height[s.numberIndex] == 0 {
					s.height[s.numberIndex] = 9
				}
			}
		} else {
			switch s.valueIndex {
			case 0:
				if rl.IsKeyPressed(state.KEY_F6) && s.lat[s.numberIndex] < 9 {
					s.lat[s.numberIndex]++
				} else if rl.IsKeyPressed(state.KEY_F7) && s.lat[s.numberIndex] > 0 {
					s.lat[s.numberIndex]--
				} else if rl.IsKeyPressed(state.KEY_F6) && s.lat[s.numberIndex] == 9 {
					s.lat[s.numberIndex] = 0
				} else if rl.IsKeyPressed(state.KEY_F7) && s.lat[s.numberIndex] == 0 {
					s.lat[s.numberIndex] = 9
				}
			case 1:
				if rl.IsKeyPressed(state.KEY_F6) && s.long[s.numberIndex] < 9 {
					s.long[s.numberIndex]++
				} else if rl.IsKeyPressed(state.KEY_F7) && s.long[s.numberIndex] > 0 {
					s.long[s.numberIndex]--
				} else if rl.IsKeyPressed(state.KEY_F6) && s.long[s.numberIndex] == 9 {
					s.long[s.numberIndex] = 0
				} else if rl.IsKeyPressed(state.KEY_F7) && s.long[s.numberIndex] == 0 {
					s.long[s.numberIndex] = 9
				}
			case 2:
				if rl.IsKeyPressed(state.KEY_F6) && s.height[s.numberIndex] < 9 {
					s.height[s.numberIndex]++
				} else if rl.IsKeyPressed(state.KEY_F7) && s.height[s.numberIndex] > 0 {
					s.height[s.numberIndex]--
				} else if rl.IsKeyPressed(state.KEY_F6) && s.height[s.numberIndex] == 9 {
					s.height[s.numberIndex] = 0
				} else if rl.IsKeyPressed(state.KEY_F7) && s.height[s.numberIndex] == 0 {
					s.height[s.numberIndex] = 9
				}
			}

		}
	}

	return s
}

func (s _state) addNavPoint() {
	s.navPointIndex += 1
	navPoints := s.navMan.NavPoints
	s.navMan.NavPoints = []navigation.MappedObject{}
	j := 0
	for i := 0; i < len(navPoints)+1; i++ {
		if i == s.navPointIndex {
			s.navMan.NavPoints = append(s.navMan.NavPoints, s.navPoint)
		} else {
			s.navMan.NavPoints = append(s.navMan.NavPoints, navPoints[j])
			j++
		}
		s.navMan.NavPoints[i].Name = "Nav Point " + strconv.FormatInt(int64(i+1), 10)
	}
}

func (s _state) removeNavPoint() {
	if len(s.navMan.NavPoints) == 0 || len(s.navMan.NavPoints) <= s.navPointIndex || s.navPointIndex == -1 {
		return
	}
	s.navMan.SelectedObject--
	navPoints := s.navMan.NavPoints
	s.navMan.NavPoints = []navigation.MappedObject{}
	for i := 0; i < len(navPoints); i++ {
		if i != s.navPointIndex {
			s.navMan.NavPoints = append(s.navMan.NavPoints, navPoints[i])
		}
	}
}

func (s _state) readNavPoint() navigation.MappedObject {
	s.navPoint.Latitude = readNumber(s.lat, s.latNegative)
	s.navPoint.Longitude = readNumber(s.long, s.longNegative)
	s.navPoint.Altitude = readNumber(s.height, false) / float32(state.GlobalOptions.Units.HightConversion)
	return s.navPoint
}

func readNumber(nArray []int, negative bool) float32 {
	n := float32(0.0)
	for i := 0; i < len(nArray); i++ {
		n += float32(nArray[i]) * float32(math.Pow(10, float64(3-i)))
	}
	if negative {
		n = n * -1
	}
	return n
}

func (s _state) Draw() {
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}
	state.DrawArrowButtonsHorizontal(0)
	state.DrawButton("ENT", 8-1)
	state.DrawButton("ADD", 9-1)
	state.DrawButton("DEL", 10-1)

	state.DrawArrowButtonsVertical(11 - 1)

	cursorHeight := state.Scale(20)
	diff := state.Scale(80)
	x := int32(rl.GetRenderWidth()/2) - state.Scale(420/2)
	y := state.Scale(2*state.SCREEN_MARGIN+state.MENU_BUTTON_HEIGHT) + cursorHeight
	rl.DrawText("Name "+s.navPoint.Name, x, y, state.Scale(30), state.COLOR_SELECT)
	drawNumberEdit("Latitude", s.lat, s.latNegative, x, y+diff, false)
	drawNumberEdit("Longitude", s.long, s.longNegative, x, y+diff*2, false)
	drawNumberEdit("Height", s.height, false, x, y+diff*3, true)
	s.drawCursor(x, y+diff, diff)
}

func (s _state) drawCursor(anchorX, anchorY int32, diff int32) {
	decimalOffset := int32(0)
	if s.numberIndex > 3 {
		decimalOffset = state.Scale(30)
	}
	x := anchorX + state.Scale(420-(30*8)) + state.Scale(30*s.numberIndex) + state.Scale(5) + decimalOffset
	y := anchorY + state.Scale(int(diff)*s.valueIndex)
	padding := state.Scale(10)
	fontSize := state.Scale(30)

	rl.SetLineWidth(state.ScaleF(3))
	a := rl.Vector2{
		X: float32(x),
		Y: float32(y - padding),
	}
	b := rl.Vector2{
		X: float32(x + state.Scale(20)),
		Y: float32(y - padding),
	}
	c := rl.Vector2{
		X: float32(x + state.Scale(10)),
		Y: float32(y - padding - state.Scale(20)),
	}
	rl.DrawTriangleLines(a, b, c, state.COLOR_SELECT)

	d := rl.Vector2{
		X: float32(x),
		Y: float32(y + padding + fontSize),
	}
	e := rl.Vector2{
		X: float32(x + state.Scale(20)),
		Y: float32(y + padding + fontSize),
	}
	f := rl.Vector2{
		X: float32(x + state.Scale(10)),
		Y: float32(y + padding + state.Scale(20) + fontSize),
	}
	rl.DrawTriangleLines(d, e, f, state.COLOR_SELECT)
}

func drawNumberEdit(label string, value []int, negative bool, anchorX, anchorY int32, extended bool) {
	fontSize := state.Scale(30)
	labelLen := rl.MeasureText(label, fontSize)
	rl.DrawText(label, anchorX, anchorY, fontSize, state.COLOR_SELECT)
	rl.DrawRectangle(anchorX+labelLen+state.Scale(10), anchorY, state.Scale(420)-labelLen-state.Scale(10), fontSize, state.COLOR_UNSELECT)
	n := ""
	j := 0
	if !extended {
		j = 1
	}
	for i := 0; i < 8; i++ {
		if i == 0 && negative && !extended {
			n = "-"
		} else if i == 4 {
			n = "."
		} else if (i >= 1 || extended) && i <= 3 || i > 4 {
			n = strconv.FormatInt(int64(value[j]), 10)
			j++
		}
		nLen := rl.MeasureText(n, fontSize)
		rl.DrawText(n, anchorX+state.Scale(420-(30*8))+state.Scale(30)*int32(i)+(15-nLen/2), anchorY, fontSize, state.COLOR_SELECT)
	}
}

func splitNumber(n float32) []int {
	str := strconv.FormatFloat(math.Abs(float64(n)), 'f', 3, 64)
	split := strings.Split(str, "")

	var numArray []int
	if math.Abs(float64(n)) < 1000 {
		numArray = append(numArray, 0)
	}
	if math.Abs(float64(n)) < 100 {
		numArray = append(numArray, 0)
	}
	if math.Abs(float64(n)) < 10 {
		numArray = append(numArray, 0)
	}
	for i := 0; i < len(split); i++ {
		if split[i] == "." {
			continue
		}
		n, err := strconv.ParseInt(split[i], 10, 64)
		if err != nil {
			panic("What")
		}
		numArray = append(numArray, int(n))
	}
	return numArray
}
