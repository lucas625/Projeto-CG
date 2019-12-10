package main

import (
	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/light"
)

func main() {
	iteractions := 3
	raysPerPixel := 100

	cam := camera.LoadJSONCamera("resources/run/json/camera.json")
	lights := light.LoadJSONLights("resources/run/json/light.json")
	objects := general.LoadJSONObjects("resources/run/json/objects.json")

	outPath := "out/pathtracing"

	// getting screen
	// camMatrix := camera.CamToWorld(&cam)
	// sc := screen.InitScreen(200, 200)
	// sc.CamToWorld = &camMatrix

	// rayCaster := raycasting.InitRayCaster(objects, &sc, &cam, &lights)

	// colorScreen := rayCaster.Run()
	// visualizer.WritePPM(*colorScreen, outPath)
}
