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
	for _, p := range objPaths {
		object := obj.ReadObj(p)
		objList = append(objList, *object)
	}

	objects := general.InitObjects("rune", objList)
	objects.ObjList[0].Color = []int{255, 0, 0}

	// getting lights
	lightPaths := []string{
		"resources/run/light.obj",
	}
	lightObj := obj.ReadObj(lightPaths[0])
	lgt := light.InitLight(100, *lightObj)
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
