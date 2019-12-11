package main

import (
	"fmt"
	// "github.com/lucas625/Projeto-CG/src/algorithms/pathtracing"
	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/light"
	// "github.com/lucas625/Projeto-CG/src/screen"
	// "github.com/lucas625/Projeto-CG/src/visualizer"
	"github.com/lucas625/Projeto-CG/src/io/obj"
)

func main() {
	// iterations := 5
	// raysPerPixel := 1000

	cam := camera.LoadJSONCamera("resources/run/json/camera.json")
	lights := light.LoadJSONLights("resources/run/json/light.json")
	objects := general.LoadJSONObjects("resources/run/json/objects.json")
	newObj := obj.ReadObj("resources/run/upperBox.obj")
	objects.ObjList = append(objects.ObjList, *newObj)
	outPath := "out/pathtracing"
	fmt.Println(cam, lights, objects)
	objects.WriteJSONObjects(outPath)
	// getting screen
	// camMatrix := camera.CamToWorld(cam)
	// sc := screen.InitScreen(600, 600)
	// sc.CamToWorld = &camMatrix

	// pathTracer := pathtracing.InitPathTracer(objects, &sc, cam, lights)

	// colorScreen := pathTracer.Run(raysPerPixel, iterations)
	// visualizer.WritePPM(*colorScreen, outPath)
}
