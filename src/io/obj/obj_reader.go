package obj

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucas625/Projeto-CG/src/general"

	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/utils"
)

func readLine(line string) (int, *[]byte) {
	result := make([]byte, 10) //fix me
	var hash int
	switch line[0] {
	case 'v':
		if line[1] == 'n' {
			hash = 1
			fmt.Println("vertice normal")
		} else {
			hash = 0
			fmt.Println("point")
		}
	case 'f':
		hash = 2
		fmt.Println("triangle")
	}
	return hash, &result
}

// readLines is a function to read allline of an .obj file.
//
// Parameters:
// 	scanner   - pointer to a Scanner with an .obj file.
//
// Returns:
//  vertices  - pointer to Vertices object with all vertices.
//  triangles - pointer to list of triangles.
//  normals   - pointer to list of normal vectors.
//
func readLines(scanner *bufio.Scanner) (*entity.Vertices, *[]entity.Triangle, *[]utils.Vector) {
	var vertList []entity.Vertex
	var triangles []entity.Triangle
	var normals []utils.Vector
	for scanner.Scan() {
		hash, result := readLine(scanner.Text())
		switch hash {
		case 0:
			var vt entity.Vertex
			err := json.Unmarshal(*result, &vt)
			utils.ShowError(err, "Unable to Unmarshal Vertex.")
			vertList = append(vertList, vt)
		case 1:
			var tri entity.Triangle
			err := json.Unmarshal(*result, &tri)
			utils.ShowError(err, "Unable to Unmarshal Triangle.")
			triangles = append(triangles, tri)
		case 2:
			var vect utils.Vector
			err := json.Unmarshal(*result, &vect)
			utils.ShowError(err, "Unable to Unmarshal Vector.")
			normals = append(normals, vect)
		}
	}
	vertices := entity.InitVertices(vertList)
	return &vertices, &triangles, &normals
}

// ReadObj is a function to read a .obj file.
//
// Parameters:
// 	objPath   - path to the .obj file.
//
// Returns:
//  cam       - the cam to the object.
//  triangles - list of triangles.
//  vertices  - Vertices object with all points.
//  normals   - list of normal vectors.
//
func ReadObj(objPath string) {
	// getting abs path
	absPath, err := filepath.Abs(objPath)
	utils.ShowError(err, "Unable to find absolute path for: "+objPath+".")

	// getting the file
	file, err := os.Open(absPath)
	utils.ShowError(err, "Unable to find "+absPath+".")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	vertices, triangles, normals := readLines(scanner)
	object := general.InitObject(vertices, triangles, normals, nil)
	fmt.Println(object)

	err = scanner.Err()
	utils.ShowError(err, "Error on reading file: "+absPath+".")

}
