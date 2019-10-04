package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/io/obj"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path to a .obj file: ")
	objPath, _ := reader.ReadString('\n')
	objPath = objPath[:len(objPath)-1] // removing \n
	switch objPath {                   //test cases
	case "-1":
		objPath = "resources/obj/simple/cone.obj"
	case "-2":
		objPath = "resources/obj/simple/cube.obj"
	case "-3":
		objPath = "resources/obj/simple/plane.obj"
	case "-4":
		objPath = "resources/obj/complex/horned_ball.obj"
	case "-5":
		objPath = "resources/obj/complex/monkey_with_cube.obj"
	case "-6":
		objPath = "resources/obj/complex/spikedball.obj"
	}

	fmt.Print("Enter the path to the output folder: ")
	outPath, _ := reader.ReadString('\n')
	outPath = outPath[:len(outPath)-1] // removing \n
	if outPath == "-1" {               // test case
		outPath = "out/"
	}

	object := obj.ReadObj(objPath)

	pt := entity.Point{Coordinates: []float64{1, 2, 3}}
	object.FindCamera(&pt)
	general.WriteJSONObject(object, outPath)
	cam := object.Camera
	cam.WriteJSONCamera(outPath)

}
