package obj

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucas625/Projeto-CG/src/utils"
)

// ReadObj is a function to read a .obj file.
//
// Parameters:
// 	objPath   - path to the .obj file.
//
// Returns:
//  cam       - the cam to the object.
//  triangles - list of triangles.
//  vertices  - Vertices object with all points.
//  normals   - List of normal vectors.
//
func ReadObj(objPath string) {
	// getting abs path
	absPath, err := filepath.Abs(objPath)
	utils.ShowError(err, "Unable to find absolute path for "+objPath)

	// getting the file
	file, err := os.Open(absPath)
	utils.ShowError(err, "Unable to find  "+absPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	utils.ShowError(err, "Error on reading file."+absPath)

}
