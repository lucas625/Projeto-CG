package screen

import (
	"errors"
	"math"
	"strconv"

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
//  d        - distance viewport to cam.
//  camWorld - the matrix camera to world.
//  px       - the additional on x (0->1)
//  py       - the additional on y (0->1)
//  fov      - field of view in degrees.
//
// Returns:
// 	a Vector.
//
func (sc *Screen) PixelToWorld(x, y int, d float64, px, py, fov float64) utils.Vector {
	if x >= sc.Height || y >= sc.Width {
		utils.ShowError(errors.New("Invalid Pixel"), "X("+strconv.Itoa(x)+") or Y("+strconv.Itoa(y)+") invalid for screen("+strconv.Itoa(sc.Height)+", "+strconv.Itoa(sc.Width)+").")
	}
	camWorld := sc.CamToWorld

	aspectRatio := float64(sc.Width) / float64(sc.Height)
	alpha := (fov / 2) * math.Pi / 180.0
	z := d

	camerax := (2*(float64(x)+px)/float64(sc.Width) - 1) * aspectRatio * math.Tan(alpha)
	cameray := (1 - 2*(float64(y)+py)/float64(sc.Height)) * math.Tan(alpha)

	v := utils.InitVector(3)

	v.Coordinates[0] = camerax
	v.Coordinates[1] = cameray
	v.Coordinates[2] = z

	vMat := utils.VectorToHomogeneousCoord(&v)

	vMatPos := utils.MultMatrix(camWorld, &vMat)
	for i := 0; i < 3; i++ {
		v.Coordinates[i] = vMatPos.Values[i][0]
	}
	vNormalized := utils.NormalizeVector(&v)

	return vNormalized
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
