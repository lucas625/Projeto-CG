package main

import (
	"github.com/lucas625/Projeto-CG/src/io/obj"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the path to a .obj file: ")
	// objPath, _ := reader.ReadString('\n')
	// objPath = objPath[:len(objPath)-1] // removing \n
	objPath := "resources/obj/simple/cube.obj"

	obj.ReadObj(objPath)
}
