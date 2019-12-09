package raycasting

import (
	"math"

	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/entity"
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
// 	lp             - pixel line index.
//  cp             - pixel column index.
//
// Returns:
// 	the colored screen painted at that position.
//
func (rcaster *RayCaster) TraceRay(coloredScreen *screen.ColoredScreen, lp, cp int) {
	screenV := rcaster.PixelScreen.PixelToWorld(lp, cp, 1.0, 0.5, 0.5)
	line := entity.Line{Start: rcaster.Cam.Pos, Director: screenV}
	color := make([]int, 3)

	closestT := math.MaxFloat64
	closestObjIdx := -1
	// intersecting objects
	for objIdx, obj := range rcaster.Objs.ObjList { // iterating through all objects
		for _, triangle := range obj.Triangles { // iterating through all triangles of an object
			points := make([]entity.Point, 3)
			for pi := 0; pi < 3; pi++ { // getting triangle points
				points[pi] = obj.Vertices.Points[triangle.Vertices[pi]]
			}
			t, _, intersected := line.IntersectTriangle(points)
			if intersected {
				p := line.FindPos(t)
				if p.Coordinates[2] >= (1+rcaster.Cam.Pos.Coordinates[2]) && t < closestT {
					closestT = t
					closestObjIdx = objIdx
				}
			}
		}
	}
	// intersecting lights
	lightClosest := false
	// for _, lgt := range rcaster.Lgts.LightList {
	// 	sphere := entity.InitSphere(lgt.LightPosition, lgt.Radius)
	// 	ts, intersected := line.IntersectSphere(sphere)
	// 	if intersected && (ts[0] >= 1 || ts[1] >= 1) {
	// 		if ts[0] <= closestT || ts[1] <= closestT {
	// 			lightClosest = true
	// 			if ts[0] <= ts[1] && ts[0] >= 1 {
	// 				closestT = ts[0]
	// 			} else {
	// 				closestT = ts[1]
	// 			}
	// 		}
	// 	}
	// }
	if !lightClosest {
		if closestObjIdx != -1 {
			color = rcaster.Objs.ObjList[closestObjIdx].Color
		}
		coloredScreen.Colors[lp][cp] = color
	} else {
		coloredScreen.Colors[lp][cp] = []int{255, 255, 255}
	}

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
	for i := 0; i < rcaster.PixelScreen.Height; i++ {
		for j := 0; j < rcaster.PixelScreen.Width; j++ {
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
