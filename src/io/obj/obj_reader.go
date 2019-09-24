package obj

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// readPoint is a function to read a vertex of an .obj file.
//
// Parameters:
// 	line   - a single .obj line.
//
// Returns:
//  the vertex as a point in json bytes.
//
func readPoint(line string) *[]byte {
	inp := strings.Split(line[2:], " ")
	pt := entity.InitPoint(3)
	for i, c := range inp {
		var err error
		pt.Coordinates[i], err = strconv.ParseFloat(c, 64)
		utils.ShowError(err, "Unable to convert to coordinate to float.")
	}
	ptAsBytes, err := json.Marshal(pt)
	utils.ShowError(err, "Unable to marshal point")
	return &ptAsBytes

}

// readLine is a function to read a line of an .obj file.
//
// Parameters:
// 	line   - a single .obj line.
//
// Returns:
//  hash   - a int that identifies the type of the return.
//  result - a Triangle, Vertex or Vector.
//
func readLine(line string) (int, *[]byte) {
	result := make([]byte, 10) //fix me
	var hash int
	switch line[0] {
	case 'v':
		if line[1] == 'n' {
			hash = 1
			inp := strings.Split(line, " ")
			fmt.Println(inp)
		} else {
			hash = 0
			result = *(readPoint(line))
		}
	case 'f':
		hash = 2
		inp := strings.Split(line[2:], " ")
		fmt.Println(inp)
	}
	return hash, &result
}

// readLines is a function to read all lines of an .obj file.
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
	var pointList []entity.Point
	var triangles []entity.Triangle
	var normals []utils.Vector
	for scanner.Scan() {
		hash, result := readLine(scanner.Text())
		print(hash, result)
		switch hash {
		case 0:
			var pt entity.Point
			err := json.Unmarshal(*result, &pt)
			utils.ShowError(err, "Unable to Unmarshal Point.")
			pointList = append(pointList, pt)
			// case 1:
			// 	var tri entity.Triangle
			// 	err := json.Unmarshal(*result, &tri)
			// 	utils.ShowError(err, "Unable to Unmarshal Triangle.")
			// 	triangles = append(triangles, tri)
			// case 2:
			// 	var vect utils.Vector
			// 	err := json.Unmarshal(*result, &vect)
			// 	utils.ShowError(err, "Unable to Unmarshal Vector.")
			// 	normals = append(normals, vect)
		}
	}
	vertices := entity.InitVertices(pointList)
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
