package screen

import (
	"errors"
	"math"
	"strconv"

	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// ColoredScreen is a class for image screen with colors.
//
// Members:
// 	Colors - the color matrix.
//
type ColoredScreen struct {
	Colors [][][]int
	Screen
}

// InitColoredScreen is a function to initialize a colored screen.
//
// Parameters:
// 	width  - the screen width.
//  height - the screen height.
//
// Returns:
// 	a colored Screen.
//
func InitColoredScreen(width, height int) ColoredScreen {
	colors := make([][][]int, height)
	for i := 0; i < height; i++ {
		colors[i] = make([][]int, width)
		for j := 0; j < width; j++ {
			colors[i][j] = make([]int, 3)
		}
	}
	return ColoredScreen{Screen: Screen{Width: width, Height: height}, Colors: colors}
}

// Screen is a class for image screen.
//
// Members:
// 	Width      - the number of x pixels on the screen.
// 	Height     - the number of y pixels on the screen.
//  CamToWorld - the matrix from cam to world.
//
type Screen struct {
	Width      int
	Height     int
	CamToWorld *utils.Matrix
}

// PixelToWorld is a function to get the position of a pixel in world coordinates.
//
// Parameters:
// 	x        - position of the pixel.
//  y        - position of the pixel.
//  d        - distance viewport to cam (if negative, considered as 1).
//  camWorld - the matrix camera to world.
//  px       - the additional on x (0->1)
//  py       - the additional on y (0->1)
//
// Returns:
// 	a Point.
//
func (sc *Screen) PixelToWorld(x, y int, d float64, px, py float64) entity.Point {
	if x >= sc.Width || y >= sc.Height {
		utils.ShowError(errors.New("Invalid Pixel"), "X("+strconv.Itoa(x)+") or Y("+strconv.Itoa(y)+") invalid for screen("+strconv.Itoa(sc.Width)+", "+strconv.Itoa(sc.Height)+").")
	}
	camWorld := sc.CamToWorld

	NDCx := (float64(x) + px) / float64(sc.Width)
	NDCy := (float64(y) + py) / float64(sc.Height)

	aspectRatio := float64(sc.Width) / float64(sc.Height)
	alpha := float64(90) // field of view
	z := 1.0
	if d > 0 {
		alpha = math.Atan(1/d) * 2
		z = d
	}

	camerax := NDCx * aspectRatio * math.Tan(alpha/2)
	cameray := NDCy * math.Tan(alpha/2)
	p := entity.InitPoint(3)
	p.Coordinates[0] = camerax
	p.Coordinates[1] = cameray
	p.Coordinates[2] = z

	pHom := entity.PointToHomogeneousCoord(&p)

	pMatPos := utils.MultMatrix(camWorld, &pHom)
	for i := 0; i < 3; i++ {
		p.Coordinates[i] = pMatPos.Values[i][0]
	}

	return p
}

// InitScreen is a function to initialize a screen.
//
// Parameters:
// 	width  - the width of the screen.
//  height - the height of the screen.
//
// Returns:
// 	a Screen.
//
func InitScreen(width, height int) Screen {
	if width < 0 || height < 0 {
		utils.ShowError(errors.New("Invalid Screen"), "width("+strconv.Itoa(width)+") or height("+strconv.Itoa(height)+") invalid for screen.")
	}
	sc := Screen{Width: width, Height: height}
	return sc
}
