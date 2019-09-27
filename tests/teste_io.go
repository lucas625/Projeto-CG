package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/io/obj"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path to a .obj file: ")
	objPath, _ := reader.ReadString('\n')
	objPath = objPath[:len(objPath)-1] // removing \n

	fmt.Print("Enter the path to the output folder: ")
	outPath, _ := reader.ReadString('\n')
	outPath = outPath[:len(outPath)-1] // removing \n

	object := obj.ReadObj(objPath)

	general.WriteJSONObject(object, outPath)
}
