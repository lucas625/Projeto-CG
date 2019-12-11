package pathtracing

import (
	"fmt"
	"math"
	"math/rand"
	"time"

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

// FindNextRay is a function to find the next line.
//
// Parameters:
//  pos       - the point.
//  obj       - the object.
//  triangIdx - the index of the triangle.
//  bCoords   - baricentric coords for each respective normal.
//
// Returns:
// 	the line
//
func (ptracer *PathTracer) FindNextRay(pos entity.Point, obj general.Object, triangleIdx int, bCoords []float64) entity.Line {
	normals := make([]utils.Vector, 3)
	for i := 0; i < 3; i++ {
		normals[i] = obj.Normals[obj.Triangles[triangleIdx].Normals[i]]
	}
	resultingNormal := utils.SumVector(&normals[0], &normals[1], bCoords[0], bCoords[1])
	resultingNormal = utils.SumVector(&resultingNormal, &normals[2], 1, bCoords[2])
	resultingNormal = utils.NormalizeVector(&resultingNormal)

	ktot := obj.TransReflection + obj.DiffuseReflection + obj.SpecularReflection
	r := 0.0 + rand.Float64()*ktot
	vector := utils.Vector{Coordinates: []float64{1.0, 1.0, 1.0}}
	if r <= obj.DiffuseReflection {
		// FIXME: calculate correct diffuse reflection
		R1 := rand.Float64()
		R2 := rand.Float64()

		phi := 1 / math.Cos(math.Sqrt(R1))
		theta := 2 * math.Pi * R2

		lightPos := ptracer.Lgts.LightList[0].LightObject.GetCenter()
		Lvector := entity.ExtractVector(&pos, &lightPos)
		Lvector = utils.NormalizeVector(&Lvector)

		constantPart := 2 * utils.DotProduct(&resultingNormal, &Lvector)
		vector = utils.SumVector(&resultingNormal, &Lvector, constantPart, -1) // R = 2N(N.L) - L
		vector.Coordinates[0] = vector.Coordinates[0] * phi
		vector.Coordinates[1] = vector.Coordinates[1] * theta
	} else if r <= obj.DiffuseReflection+obj.SpecularReflection {
		lightPos := ptracer.Lgts.LightList[0].LightObject.GetCenter()
		Lvector := entity.ExtractVector(&pos, &lightPos)
		Lvector = utils.NormalizeVector(&Lvector)

		constantPart := 2 * utils.DotProduct(&resultingNormal, &Lvector)
		vector = utils.SumVector(&resultingNormal, &Lvector, constantPart, -1) // R = 2N(N.L) - L
	} else {
		// use transmission (unavailable)
	}
	vector = utils.NormalizeVector(&vector)
	line := entity.Line{Start: pos, Director: vector}
	return line
}

// TraceRayDepth is a function to trace a ray and return the resulting color.
//
// Parameters:
//  line       - the ray.
//  recursions - number of recursions.
//
// Returns:
// 	the rgb color at a given position.
//
func (ptracer *PathTracer) TraceRayDepth(line entity.Line, recursions int) []float64 {
	color := make([]float64, 3)

	closestT := math.MaxFloat64
	closestObjIdx := -1
	closestTriangleIndex := -1
	closestBCoords := make([]float64, 3)
	for objIdx, obj := range ptracer.Objs.ObjList { // iterating through all objects
		for triangIdx, triangle := range obj.Triangles { // iterating through all triangles of an object
			points := make([]entity.Point, 3)
			for pi := 0; pi < 3; pi++ { // getting triangle points
				points[pi] = obj.Vertices.Points[triangle.Vertices[pi]]
			}
			t, bCoords, intersected := line.IntersectTriangle(points)
			if intersected {
				if t > 0 && t < closestT {
					closestT = t
					closestObjIdx = objIdx
					closestTriangleIndex = triangIdx
					closestBCoords = bCoords
				}
			}
		}
	}

	// intersecting lights
	lightClosest := false
	for lgtIdx, lgt := range ptracer.Lgts.LightList {
		for _, triangle := range lgt.LightObject.Triangles { // iterating through all triangles of an object
			points := make([]entity.Point, 3)
			for pi := 0; pi < 3; pi++ { // getting triangle points
				points[pi] = lgt.LightObject.Vertices.Points[triangle.Vertices[pi]]
			}
			t, _, intersected := line.IntersectTriangle(points)
			if intersected {
				if t > 0 && t <= closestT {
					lightClosest = true
					closestT = t
					closestObjIdx = lgtIdx
				}
			}
		}
	}

	if !lightClosest {
		if closestObjIdx != -1 {
			colorAux := []float64{1, 1, 1}
			if recursions > 0 {
				newLine := ptracer.FindNextRay(line.FindPos(closestT), ptracer.Objs.ObjList[closestObjIdx], closestTriangleIndex, closestBCoords)
				colorAux = ptracer.TraceRayDepth(newLine, recursions-1)
			}
			for i := 0; i < 3; i++ {
				color[i] = ptracer.Objs.ObjList[closestObjIdx].Color[i] * colorAux[i]
			}
		}
	} else {
		for i := 0; i < 3; i++ {
			lgtAux := ptracer.Lgts.LightList[closestObjIdx]
			color[i] = lgtAux.Color[i] * lgtAux.LightIntensity
		}
	}
	return color
}

// TraceRay is a function to trace a ray through a pixel.
//
// Parameters:
// 	lp             - pixel line index.
//  cp             - pixel column index.
//  rays           - number of rays per pixel.
//  recursions     - number of recursions.
//
// Returns:
// 	the colored screen painted at that position.
//
func (ptracer *PathTracer) TraceRay(lp, cp, rays, recursions int) []int {
	floatColors := make([][]float64, rays)
	for ray := 0; ray < rays; ray++ {

		offx := rand.Float64()
		offy := rand.Float64()

		screenV := ptracer.PixelScreen.PixelToWorld(lp, cp, 1.0, offx, offy, ptracer.Cam.FieldOfView)
		line := entity.Line{Start: ptracer.Cam.Pos, Director: screenV}

		color := make([]float64, 3)

		closestT := math.MaxFloat64
		closestObjIdx := -1
		closestTriangleIndex := -1
		closestBCoords := make([]float64, 3)
		for objIdx, obj := range ptracer.Objs.ObjList { // iterating through all objects
			for triangIdx, triangle := range obj.Triangles { // iterating through all triangles of an object
				points := make([]entity.Point, 3)
				for pi := 0; pi < 3; pi++ { // getting triangle points
					points[pi] = obj.Vertices.Points[triangle.Vertices[pi]]
				}
				t, bCoords, intersected := line.IntersectTriangle(points)
				if intersected {
					if t >= 1 && t < closestT {
						closestT = t
						closestObjIdx = objIdx
						closestTriangleIndex = triangIdx
						closestBCoords = bCoords
					}
				}
			}
		}

		// intersecting lights
		lightClosest := false
		for lgtIdx, lgt := range ptracer.Lgts.LightList {
			for _, triangle := range lgt.LightObject.Triangles { // iterating through all triangles of an object
				points := make([]entity.Point, 3)
				for pi := 0; pi < 3; pi++ { // getting triangle points
					points[pi] = lgt.LightObject.Vertices.Points[triangle.Vertices[pi]]
				}
				t, _, intersected := line.IntersectTriangle(points)
				if intersected {
					if t >= 1 && t <= closestT {
						lightClosest = true
						closestT = t
						closestObjIdx = lgtIdx
					}
				}
			}
		}

		if !lightClosest {
			if closestObjIdx != -1 {
				colorAux := []float64{1, 1, 1}
				if recursions > 0 {
					newLine := ptracer.FindNextRay(line.FindPos(closestT), ptracer.Objs.ObjList[closestObjIdx], closestTriangleIndex, closestBCoords)
					colorAux = ptracer.TraceRayDepth(newLine, recursions-1)
				}
				for i := 0; i < 3; i++ {
					color[i] = ptracer.Objs.ObjList[closestObjIdx].Color[i] * colorAux[i]
				}
			}
		} else {
			for i := 0; i < 3; i++ {
				lgtAux := ptracer.Lgts.LightList[closestObjIdx]
				color[i] = lgtAux.Color[i] * lgtAux.LightIntensity
			}
		}
		floatColors[ray] = color
	}

	// calculating average
	color := make([]float64, 3)
	for i := 0; i < rays; i++ {
		for j := 0; j < 3; j++ {
			color[j] = color[j] + (floatColors[i][j] / float64(rays))
		}
	}

	intColor := make([]int, 3)
	for i := 0; i < 3; i++ {
		intColor[i] = int(math.Floor(math.Sqrt(color[i]) * 255))
		if intColor[i] > 255 {
			intColor[i] = 255
		} else if intColor[i] < 0 {
			intColor[i] = 0
		}
	}
	return intColor

}

// Run is a function to run the ray casting.
//
// Parameters:
// 	none
//
// Returns:
// 	the colored screen.
//
func (ptracer *PathTracer) Run(rays, recursions int) *screen.ColoredScreen {
	coloredScreen := screen.InitColoredScreen(ptracer.PixelScreen.Width, ptracer.PixelScreen.Height)
	for i := 0; i < ptracer.PixelScreen.Height; i++ {
		for j := 0; j < ptracer.PixelScreen.Width; j++ {
			fmt.Println(i, j)
			coloredScreen.Colors[i][j] = ptracer.TraceRay(i, j, rays, recursions)
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
	rand.Seed(time.Now().UnixNano())
	return PathTracer{Objs: objs, PixelScreen: pixelScreen, Cam: cam, Lgts: lgts}
}
