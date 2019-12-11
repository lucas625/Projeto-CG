package main

import (
	"fmt"

	"github.com/lucas625/Projeto-CG/src/algorithms/pathtracing"
	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/light"
	"github.com/lucas625/Projeto-CG/src/screen"
	"github.com/lucas625/Projeto-CG/src/visualizer"
)

func main() {
	iterations := 10
	raysPerPixel := 500

	cam := camera.LoadJSONCamera("resources/run/json/camera.json")
	lights := light.LoadJSONLights("resources/run/json/light.json")
	objects := general.LoadJSONObjects("resources/run/json/objects.json")
	fmt.Println(cam)
	fmt.Println(lights)
	fmt.Println(objects)
	outPath := "out/pathtracing"

	// getting screen
	camMatrix := camera.CamToWorld(cam)
	sc := screen.InitScreen(200, 200)
	sc.CamToWorld = &camMatrix

	pathTracer := pathtracing.InitPathTracer(objects, &sc, cam, lights)

	colorScreen := pathTracer.Run(raysPerPixel, iterations)
	visualizer.WritePPM(*colorScreen, outPath)
}
