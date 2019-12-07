package raycasting

import (
	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/light"
	"github.com/lucas625/Projeto-CG/src/screen"
)

// RayCaster is a class for ray casting algorithm.
//
// Members:
// 	Objs        - the list of objects.
//  PixelScreen - the screen.
//  Cam         - the camera.
//  Lgts        - the lights.
//
type RayCaster struct {
	Objs        *general.Objects
	PixelScreen *screen.Screen
	Cam         *camera.Camera
	Lgts        *light.Lights
}

// TraceRay is a function to trace a ray through a pixel.
//
// Parameters:
//  coloredScreen - the screen to be painted.
// 	i             - pixel line index.
//  j             - pixel column index.
//
// Returns:
// 	the colored screen painted at that position.
//
func (rcaster *RayCaster) TraceRay(coloredScreen *screen.ColoredScreen, i, j int) {

}

// Run is a function to run the ray casting.
//
// Parameters:
// 	none
//
// Returns:
// 	the colored screen.
//
func (rcaster *RayCaster) Run() *screen.ColoredScreen {
	coloredScreen := screen.InitColoredScreen(rcaster.PixelScreen.Width, rcaster.PixelScreen.Height)
	for i := 0; i < rcaster.PixelScreen.Width; i++ {
		for j := 0; j < rcaster.PixelScreen.Height; j++ {
			rcaster.TraceRay(&coloredScreen, i, j)
		}
	}
	return &coloredScreen
}

// InitRayCaster is a function to initialize a RayCaster.
//
// Parameters:
// 	objs        - the list of objects.
//  pixelScreen - the screen.
//  cam         - the camera.
//  lgts        - the lights.
//
// Returns:
// 	a RayCaster.
//
func InitRayCaster(objs *general.Objects, pixelScreen *screen.Screen, cam *camera.Camera, lgts *light.Lights) RayCaster {
	return RayCaster{Objs: objs, PixelScreen: pixelScreen, Cam: cam, Lgts: lgts}
}
