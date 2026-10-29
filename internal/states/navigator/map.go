package navigator

import (
	"fmt"
	"math"
	"strconv"

	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
 * Map goes from **Bottom-Rigth** to **Top-Left** row by row.
 * The map has 16 images for every row and 8 for every column
 * Images have a 1000x1000 resolution displaying a RESOLUTION°xRESOLUTION°
 * grid area.
 * Starting at `map-1.jpg` and ending at `map-128.jpg`.
 */
const RESOLUTION float64 = 2.25

type gridImage struct {
	columnIndex int // Index inside a column 0-7
	rowIndex    int // Index inside a row 0-15
	image       rl.Texture2D
}

type mapGrid struct {
	center int // Which of 0,1,2,3 is the center of the map.
	images []gridImage
}

func (s _state) DrawMap(anchorX, anchorY int32, width int32) {
	lats, longs := s.calcGridCoords() // 3x3 Grid of coordinates

	/* colors := []rl.Color{
		rl.Blue,
		rl.Green,
		rl.Red,
		rl.Yellow,
	} */

	xs := []float32{}
	ys := []float32{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			x, y, _, _ := s.navMan.CoordToDistance(
				float64(state.GlobalFlightData.Latitude),
				float64(state.GlobalFlightData.Longitude),
				lats[i],
				longs[j],
				float64(state.GlobalFlightData.Heading),
			)
			relX := (x * state.GlobalOptions.Units.DistanceConversion / float64(s.scale)) * float64(width/4) * 3
			relY := (y * state.GlobalOptions.Units.DistanceConversion / float64(s.scale)) * float64(width/4) * 3

			xs = append(xs, float32(anchorX)+float32(relX))
			ys = append(ys, float32(anchorY)+float32(relY))
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			vecs := []rl.Vector2{
				{X: xs[i*3+j], Y: ys[i*3+j]},
				{X: xs[(i+1)*3+j], Y: ys[(i+1)*3+j]},
				{X: xs[i*3+(j+1)], Y: ys[i*3+(j+1)]},
				{X: xs[(i+1)*3+(j+1)], Y: ys[(i+1)*3+(j+1)]},
			}
			tex := s.mapGrid.images[j*2+i].image
			state.DrawQuadTexture(vecs, tex)
			// col := colors[i*2+j]
			// rl.DrawRing(rl.Vector2{X: xs[i*3+j], Y: ys[i*3+j]}, 4, 8, 0, 360, 0, col)
			// state.DrawQuadColor(vecs, col)
		}
	}
}

func (s _state) calcGridCoords() ([]float64, []float64) {
	var closestPointLat float64
	var closestPointLong float64
	if (math.Mod(float64(state.GlobalFlightData.Latitude)+90, RESOLUTION)) < RESOLUTION/2 {
		closestPointLat = float64(state.GlobalFlightData.Latitude) - (math.Mod(float64(state.GlobalFlightData.Latitude)+90, RESOLUTION))
	} else {
		closestPointLat = float64(state.GlobalFlightData.Latitude) - (math.Mod(float64(state.GlobalFlightData.Latitude)+90, RESOLUTION)) + RESOLUTION
	}
	if (math.Mod(float64(state.GlobalFlightData.Longitude)+360, RESOLUTION)) < RESOLUTION/2 {
		closestPointLong = float64(state.GlobalFlightData.Longitude) - (math.Mod(float64(state.GlobalFlightData.Longitude)+360, RESOLUTION))
	} else {
		closestPointLong = float64(state.GlobalFlightData.Longitude) - (math.Mod(float64(state.GlobalFlightData.Longitude)+360, RESOLUTION)) + RESOLUTION
	}
	// fmt.Printf("Closest Point (%f, %f)\n", closestPointLat, closestPointLong)
	lats := []float64{closestPointLat + RESOLUTION, closestPointLat, closestPointLat - RESOLUTION}
	longs := []float64{closestPointLong + RESOLUTION, closestPointLong, closestPointLong - RESOLUTION}

	return lats, longs
}

func (s _state) calcMap() (_state, error) {
	var closestPointLat float64
	var closestPointLong float64
	if (math.Mod(float64(state.GlobalFlightData.Latitude)+90, RESOLUTION)) < RESOLUTION/2 {
		closestPointLat = float64(state.GlobalFlightData.Latitude) - (math.Mod(float64(state.GlobalFlightData.Latitude)+90, RESOLUTION))
	} else {
		closestPointLat = float64(state.GlobalFlightData.Latitude) - (math.Mod(float64(state.GlobalFlightData.Latitude)+90, RESOLUTION)) + RESOLUTION
	}
	if (math.Mod(float64(state.GlobalFlightData.Longitude)+360, RESOLUTION)) < RESOLUTION/2 {
		closestPointLong = float64(state.GlobalFlightData.Longitude) - (math.Mod(float64(state.GlobalFlightData.Longitude)+360, RESOLUTION))
	} else {
		closestPointLong = float64(state.GlobalFlightData.Longitude) - (math.Mod(float64(state.GlobalFlightData.Longitude)+360, RESOLUTION)) + RESOLUTION
	}
	gridColumnIndex := int(180/RESOLUTION) - int((closestPointLat+90)/RESOLUTION)
	gridRowIndex := int((closestPointLong + 180) / RESOLUTION)

	if len(s.mapGrid.images) != 4 {
		null := gridImage{
			columnIndex: 0,
			rowIndex:    0,
			image:       rl.Texture2D{},
		}
		s.mapGrid.images = []gridImage{null, null, null, null}
		fmt.Printf("Filled Image Array\n")
	}

	var err error
	s.mapGrid.images[0], err = s.loadMapImage(gridColumnIndex-1, gridRowIndex, 0)
	if err != nil {
		return s, err
	}
	s.mapGrid.images[1], err = s.loadMapImage(gridColumnIndex, gridRowIndex, 1)
	if err != nil {
		return s, err
	}
	s.mapGrid.images[2], err = s.loadMapImage(gridColumnIndex-1, gridRowIndex-1, 2)
	if err != nil {
		return s, err
	}
	s.mapGrid.images[3], err = s.loadMapImage(gridColumnIndex, gridRowIndex-1, 3)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s _state) loadMapImage(columnIndex int, rowIndex int, imgIndex int) (gridImage, error) {
	if columnIndex < 0 {
		columnIndex = int(180/RESOLUTION) + columnIndex
	}
	if rowIndex < 0 {
		rowIndex = int(360/RESOLUTION) + rowIndex
	}
	columnIndex = columnIndex % int(180/RESOLUTION)
	rowIndex = rowIndex % int(360/RESOLUTION)
	var img rl.Texture2D
	// num := columnIndex*int(360/RESOLUTION) + rowIndex

	if columnIndex == s.mapGrid.images[imgIndex].columnIndex && rowIndex == s.mapGrid.images[imgIndex].rowIndex {
		img = s.mapGrid.images[imgIndex].image
		// fmt.Printf("Same Image, not reloading.\n")
	} else {
		rl.UnloadTexture(s.mapGrid.images[imgIndex].image)
		// img = rl.LoadTexture("assets/map/map-" + strconv.FormatInt(int64(num), 10) + ".jpg")
		img = rl.LoadTexture("assets/map/" + strconv.FormatInt(int64(columnIndex), 10) + "_" + strconv.FormatInt(int64(rowIndex), 10) + ".png")
	}

	return gridImage{
		columnIndex: columnIndex % int(180/RESOLUTION),
		rowIndex:    rowIndex % int(360/RESOLUTION),
		image:       img,
	}, nil
}

func (s _state) DrawMapMargin() {
	rl.DrawRectangle(0, 0, state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), int32(rl.GetRenderHeight()), rl.Black)
	rl.DrawRectangle(int32(rl.GetRenderWidth())-state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), 0, state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), int32(rl.GetRenderHeight()), rl.Black)
	rl.DrawRectangle(0, 0, int32(rl.GetRenderWidth()), state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), rl.Black)
	rl.DrawRectangle(0, int32(rl.GetRenderHeight())-state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), int32(rl.GetRenderWidth()), state.Scale(state.MENU_BUTTON_WIDTH+state.SCREEN_MARGIN*2), rl.Black)

}
