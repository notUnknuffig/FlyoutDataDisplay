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

func (s _state) loadMapImage(columnIndex int, rowIndex int, imgIndex int) gridImage {
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
		fmt.Printf("Same Image, not reloading.\n")
	} else {
		rl.UnloadTexture(s.mapGrid.images[imgIndex].image)
		// img = rl.LoadTexture("assets/map/map-" + strconv.FormatInt(int64(num), 10) + ".jpg")
		img = rl.LoadTexture("assets/map/" + strconv.FormatInt(int64(columnIndex), 10) + "_" + strconv.FormatInt(int64(rowIndex), 10) + ".jpg")
	}

	return gridImage{
		columnIndex: columnIndex % int(180/RESOLUTION),
		rowIndex:    rowIndex % int(360/RESOLUTION),
		image:       img,
	}
}

func (s _state) calcMap() _state {
	gridColumnIndex := int(math.Floor(float64(state.GlobalFlightData.Latitude+90) / RESOLUTION)) // 0 - 7
	gridRowIndex := int(math.Floor(float64(state.GlobalFlightData.Longitude+180) / RESOLUTION))  // 1 - 16
	imageIndex := gridColumnIndex*16 + gridRowIndex + 1

	inGridLatitude := math.Mod(float64(state.GlobalFlightData.Latitude+90), RESOLUTION)
	inGridLongitude := math.Mod(float64(state.GlobalFlightData.Longitude+180), RESOLUTION)
	fmt.Printf(
		"Coord (%f, %f), Index (%d, %d) Image %d, In grid coord (%f, %f)\n",
		state.GlobalFlightData.Latitude,
		state.GlobalFlightData.Longitude,
		gridColumnIndex,
		gridRowIndex,
		imageIndex,
		inGridLatitude,
		inGridLongitude,
	)

	fmt.Printf("Array: %d\n", s.mapGrid.images)
	if len(s.mapGrid.images) != 4 {
		null := gridImage{
			columnIndex: 8,
			rowIndex:    16,
			image:       rl.Texture2D{},
		}
		s.mapGrid.images = []gridImage{null, null, null, null}
		fmt.Printf("Filled Image Array\n")
	}

	if inGridLatitude >= 11.25 {
		// Top
		if inGridLongitude >= 11.25 {
			// Top Right -> Bottom Left
			s.mapGrid.center = 3
			s.mapGrid.images[0] = s.loadMapImage(gridColumnIndex+1, gridRowIndex+1, 0)
			s.mapGrid.images[1] = s.loadMapImage(gridColumnIndex+1, gridRowIndex, 1)
			s.mapGrid.images[2] = s.loadMapImage(gridColumnIndex, gridRowIndex+1, 2)
			s.mapGrid.images[3] = s.loadMapImage(gridColumnIndex, gridRowIndex, 3)
		} else if inGridLongitude < 11.25 {
			// Top Left -> Bottom Right
			s.mapGrid.center = 2
			s.mapGrid.images[0] = s.loadMapImage(gridColumnIndex+1, gridRowIndex, 0)
			s.mapGrid.images[1] = s.loadMapImage(gridColumnIndex+1, gridRowIndex-1, 1)
			s.mapGrid.images[2] = s.loadMapImage(gridColumnIndex, gridRowIndex, 2)
			s.mapGrid.images[3] = s.loadMapImage(gridColumnIndex, gridRowIndex-1, 3)
		}
	} else if inGridLatitude < 11.25 {
		// Bottom
		if inGridLongitude >= 11.25 {
			// Bottom Right -> Top Left
			s.mapGrid.center = 1
			s.mapGrid.images[0] = s.loadMapImage(gridColumnIndex, gridRowIndex+1, 0)
			s.mapGrid.images[1] = s.loadMapImage(gridColumnIndex, gridRowIndex, 1)
			s.mapGrid.images[2] = s.loadMapImage(gridColumnIndex-1, gridRowIndex+1, 2)
			s.mapGrid.images[3] = s.loadMapImage(gridColumnIndex-1, gridRowIndex, 3)
		} else if inGridLongitude < 11.25 {
			// Bottom Left -> Top Right
			s.mapGrid.center = 0
			s.mapGrid.images[0] = s.loadMapImage(gridColumnIndex, gridRowIndex, 0)
			s.mapGrid.images[1] = s.loadMapImage(gridColumnIndex, gridRowIndex-1, 1)
			s.mapGrid.images[2] = s.loadMapImage(gridColumnIndex-1, gridRowIndex, 2)
			s.mapGrid.images[3] = s.loadMapImage(gridColumnIndex-1, gridRowIndex-1, 3)
		}
	}
	return s
}

func (s _state) DrawMap() {
	inGridLatitude := math.Mod(float64(state.GlobalFlightData.Latitude+90), RESOLUTION) / RESOLUTION
	inGridLongitude := math.Mod(float64(state.GlobalFlightData.Longitude+180), RESOLUTION) / RESOLUTION
	for i := 0; i < len(s.mapGrid.images); i++ {
		y := 1
		if i < 2 {
			y = 0
		}
		rl.DrawTextureEx(s.mapGrid.images[i].image, rl.Vector2{X: 400*float32(i%2) + 400*float32(inGridLongitude), Y: 400*float32(y) + 400*float32(inGridLatitude)}, 0, 0.4, rl.White)
	}
}

func (s _state) DrawMapMargin() {
	rl.DrawRectangle(0, 0, state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), int32(rl.GetRenderHeight()), rl.Black)
	rl.DrawRectangle(int32(rl.GetRenderWidth())-state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), 0, state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), int32(rl.GetRenderHeight()), rl.Black)
	rl.DrawRectangle(0, 0, int32(rl.GetRenderWidth()), state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), rl.Black)
	rl.DrawRectangle(0, int32(rl.GetRenderHeight())-state.Scale(state.MENU_BUTTON_HEIGHT+state.SCREEN_MARGIN*2), int32(rl.GetRenderWidth()), state.Scale(state.MENU_BUTTON_WIDTH+state.SCREEN_MARGIN*2), rl.Black)

}
