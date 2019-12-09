package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/lucas625/Projeto-CG/src/utils"

	"github.com/lucas625/Projeto-CG/src/algorithms/raycasting"
	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/io/obj"
	"github.com/lucas625/Projeto-CG/src/light"
	"github.com/lucas625/Projeto-CG/src/screen"
	"github.com/lucas625/Projeto-CG/src/visualizer"
)

func main() {
	ended := false
	objList := make([]general.Object, 0, 10)
	current := 0
	reader := bufio.NewReader(os.Stdin)
	for !ended {
		fmt.Print("Enter the path to a .obj file: ")
		objPath, _ := reader.ReadString('\n')
		obi, err := strconv.Atoi(objPath[:len(objPath)-2]) // removing \n
		utils.ShowError(err, "fail")
		switch obi { //test cases
		case 1:
			objPath = "resources/obj/simple/cone.obj"
		case 2:
			objPath = "resources/obj/simple/cube.obj"
		case 3:
			objPath = "resources/obj/simple/plane.obj"
		case 4:
			objPath = "resources/obj/complex/horned_ball.obj"
		case 5:
			objPath = "resources/obj/complex/monkey_with_cube.obj"
		case 6:
			objPath = "resources/obj/complex/spikedball.obj"
		case 0:
			if current >= 1 {
				ended = true
				continue
			} else {
				utils.ShowError(errors.New("Need at least one object"), "Try again")
			}
		}

		object := obj.ReadObj(objPath)
		objList = append(objList, *object)
		current++
	}

	fmt.Print("Enter the path to the output folder: ")
	outPath, _ := reader.ReadString('\n')
	outPath = outPath[:len(outPath)-2] // removing \n
	if outPath == "-1" {               // test case
		outPath = "out/"
	}

	objects := general.InitObjects("teste", objList)
	objects.ObjList[0].Color = []int{255, 0, 0}

	cameraPath := "resources/json/camera.json"
	cam := camera.LoadJSONCamera(cameraPath)
	// cam = objects.ObjList[0].FindCamera(&cam.Pos)

	camMatrix := camera.CamToWorld(cam)

	sc := screen.InitScreen(200, 200)
	sc.CamToWorld = &camMatrix

	lightPath := "resources/json/light.json"
	lights := light.LoadJSONLights(lightPath)

	rayCaster := raycasting.InitRayCaster(objects, &sc, cam, lights)

	colorScreen := rayCaster.Run()
	visualizer.WritePPM(*colorScreen, outPath)
}
