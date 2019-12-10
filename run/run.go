package main

import (
	"github.com/lucas625/Projeto-CG/src/algorithms/raycasting"
	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/io/obj"
	"github.com/lucas625/Projeto-CG/src/light"
	"github.com/lucas625/Projeto-CG/src/screen"
	"github.com/lucas625/Projeto-CG/src/visualizer"
)

func main() {
	// ---------- PARAMETERS ----------
	iteractions := 3
	raysPerPixel := 100
	specularDecayList := []float64{
		5,
		5,
		5,
		5,
		5,
		5,
	}
	kaList := []float64{
		0.3,
		0.3,
		0.3,
		0.3,
		0.3,
		0.3,
	}
	kdList := []float64{
		0.7,
		0.7,
		0.7,
		0.7,
		0.7,
		0.7,
	}
	ksList := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
	}
	ktList := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
	}

	// getting objects
	objList := make([]general.Object, 0, 10)
	objPaths := []string{
		"resources/run/back.obj",
		"resources/run/left_wall.obj",
		"resources/run/right_wall.obj",
		"resources/run/ground.obj",
		"resources/run/ceiling.obj",
		"resources/run/box-only.obj",
	}
	for i, p := range objPaths {
		object := obj.ReadObj(p)
		object.AmbientReflection = kaList[i]
		object.DiffuseReflection = kdList[i]
		object.SpecularReflection = ksList[i]
		object.TransReflection = ktList[i]
		object.SpecularDecay = specularDecayList[i]
		objList = append(objList, *object)
	}

	objects := general.InitObjects("rune", objList)

	// getting lights
	lightPaths := []string{
		"resources/run/light.obj",
	}
	lightObj := obj.ReadObj(lightPaths[0])
	lgt := light.InitLight(
		make([]int, 3),
		[]int{255, 255, 255},
		*lightObj)
	lights := light.Lights{LightList: []light.Light{lgt}}

	outPath := "out/pathtracing/"

	// getting camera
	cameraPath := "resources/json/camera.json"
	camAux := camera.LoadJSONCamera(cameraPath)
	pAux := entity.InitPoint(3)
	for i := 0; i < 3; i++ {
		pAux.Coordinates[i] = camAux.Pos.Coordinates[i] + camAux.Look.Coordinates[i]
	}
	cam := camera.InitCameraWithPoints(&camAux.Pos, &pAux)

	// getting screen
	camMatrix := camera.CamToWorld(&cam)
	sc := screen.InitScreen(200, 200)
	sc.CamToWorld = &camMatrix

	rayCaster := raycasting.InitRayCaster(objects, &sc, &cam, &lights)

	colorScreen := rayCaster.Run()
	visualizer.WritePPM(*colorScreen, outPath)
}
