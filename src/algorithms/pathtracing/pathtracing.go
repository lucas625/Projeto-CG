package raycasting

import (
	"errors"
	"math"

	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/light"
	"github.com/lucas625/Projeto-CG/src/screen"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// PathTracer is a class for path tracing algorithm.
//
// Members:
// 	Objs        - the list of objects.
//  PixelScreen - the screen.
//  Cam         - the camera.
//  Lgts        - the lights.
//
type PathTracer struct {
	Objs        *general.Objects
	PixelScreen *screen.Screen
	Cam         *camera.Camera
	Lgts        *light.Lights
}

// CalculateColor is a function to calculate the color at a point.
//
// Parameters:
//
//
// Returns:
// 	the color at a given point.
//
func (ptracer *PathTracer) CalculateColor(objIdx, triangleIdx int, bcoords []float64, pos entity.Point) []int {
	resultColor := make([]int, 3)
	obj := ptracer.Objs.ObjList[objIdx]
	Nvector := obj.GetNormalByBaricentricCoords(triangleIdx, bcoords)
	Lvector := utils.InitVector(3) // FIXME: find the real L vector.
	raux := utils.CMultVector(&Nvector, utils.DotProduct(&Nvector, &Lvector)*2)
	Rvector := utils.SumVector(&raux, &Lvector, 1, -1)
	Vvector := entity.ExtractVector(&pos, &ptracer.Cam.Pos)
	for _, lgt := range ptracer.Lgts.LightList {
		ambientColor := make([]float64, 3)
		diffuseColor := make([]float64, 3)
		specularColor := make([]float64, 3)
		for i := 0; i < 3; i++ {
			ambientColor[i] = lgt.AmbientIntensity * obj.AmbientReflection * float64(obj.Color[i])
			diffuseColor[i] = lgt.LightIntensity * obj.DiffuseReflection * float64(obj.Color[i])
			specularColor[i] = lgt.LightIntensity * obj.SpecularReflection * math.Pow(utils.DotProduct(&Rvector, &Vvector), obj.SpecularDecay) * float64(lgt.Color[i])
			resultColor[i] = resultColor[i] + int(math.Floor(ambientColor[i]+diffuseColor[i]+specularColor[i]))
		}
	}
	for i := 0; i < 3; i++ {
		if resultColor[i] > 255 {
			resultColor[i] = 255
		} else if resultColor[i] < 0 {
			utils.ShowError(errors.New("Invalid color"), "Color less than 0")
		}
	}
	return resultColor
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
func (ptracer *PathTracer) TraceRay(coloredScreen *screen.ColoredScreen, lp, cp int) {
	screenV := ptracer.PixelScreen.PixelToWorld(lp, cp, 1.0, 0.5, 0.5)
	line := entity.Line{Start: ptracer.Cam.Pos, Director: screenV}
	color := make([]int, 3)

	closestT := math.MaxFloat64
	closestObjIdx := -1
	// intersecting objects
	for objIdx, obj := range ptracer.Objs.ObjList { // iterating through all objects
		for _, triangle := range obj.Triangles { // iterating through all triangles of an object
			points := make([]entity.Point, 3)
			for pi := 0; pi < 3; pi++ { // getting triangle points
				points[pi] = obj.Vertices.Points[triangle.Vertices[pi]]
			}
			t, _, intersected := line.IntersectTriangle(points)
			if intersected {
				p := line.FindPos(t)
				if p.Coordinates[2] >= (1+ptracer.Cam.Pos.Coordinates[2]) && t < closestT {
					closestT = t
					closestObjIdx = objIdx
				}
			}
		}
	}
	// intersecting lights
	lightClosest := false
	// for _, lgt := range ptracer.Lgts.LightList {
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
			color = ptracer.Objs.ObjList[closestObjIdx].Color
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
func (ptracer *PathTracer) Run() *screen.ColoredScreen {
	coloredScreen := screen.InitColoredScreen(ptracer.PixelScreen.Width, ptracer.PixelScreen.Height)
	for i := 0; i < ptracer.PixelScreen.Height; i++ {
		for j := 0; j < ptracer.PixelScreen.Width; j++ {
			ptracer.TraceRay(&coloredScreen, i, j)
		}
	}
	return &coloredScreen
}

// InitPathTracer is a function to initialize a PathTracer.
//
// Parameters:
// 	objs        - the list of objects.
//  pixelScreen - the screen.
//  cam         - the camera.
//  lgts        - the lights.
//
// Returns:
// 	a PathTracer.
//
func InitPathTracer(objs *general.Objects, pixelScreen *screen.Screen, cam *camera.Camera, lgts *light.Lights) PathTracer {
	return PathTracer{Objs: objs, PixelScreen: pixelScreen, Cam: cam, Lgts: lgts}
}
