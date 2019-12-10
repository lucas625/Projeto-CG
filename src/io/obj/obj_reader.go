package obj

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// readTrianle is a function to read a face of an .obj file.
//
// Parameters:
// 	line   - a single .obj line.
//
// Returns:
//  the triangle as json bytes.
//
func readTriangle(line string) *[]byte {
	inp := strings.Split(line[2:], " ")
	vertices := make([]int, 3)
	normals := make([]int, 3)
	for i, c := range inp {
		splitedIndices := strings.Split(c, "/")
		var err error
		vertices[i], err = strconv.Atoi(splitedIndices[0])
		utils.ShowError(err, "Unable to convert to vertex index to int.")
		normals[i], err = strconv.Atoi(splitedIndices[len(splitedIndices)-1])
		utils.ShowError(err, "Unable to convert to normal index to int.")
		vertices[i] = vertices[i] - 1
		normals[i] = normals[i] - 1
	}
	triangle := entity.InitTriangle(vertices, normals)
	triangleAsBytes, err := json.Marshal(triangle)
	utils.ShowError(err, "Unable to marshal triangle.")
	return &triangleAsBytes
}

// readNormal is a function to read a vertex normal of an .obj file.
//
// Parameters:
// 	line   - a single .obj line.
//
// Returns:
//  the vector as json bytes.
//
func readNormal(line string) *[]byte {
	inp := strings.Split(line[3:], " ")
	vect := utils.InitVector(3)
	for i, c := range inp {
		var err error
		vect.Coordinates[i], err = strconv.ParseFloat(c, 64)
		utils.ShowError(err, "Unable to convert coordinate to float.")
	}
	vectAsBytes, err := json.Marshal(vect)
	utils.ShowError(err, "Unable to marshal vector.")
	return &vectAsBytes
}

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
		utils.ShowError(err, "Unable to convert coordinate to float.")
	}
	ptAsBytes, err := json.Marshal(pt)
	utils.ShowError(err, "Unable to marshal point.")
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
	switch line[0] {
	case 'v':
		if line[1] == 'n' {
			return 2, readNormal(line)
		}
		return 0, readPoint(line) // vertex case
	case 'f':
		return 1, readTriangle(line)
	}
	return -1, nil
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
func readLines(scanner *bufio.Scanner) (entity.Vertices, []entity.Triangle, []utils.Vector) {
	var pointList []entity.Point
	var triangles []entity.Triangle
	var normals []utils.Vector
	for scanner.Scan() {
		hash, result := readLine(scanner.Text())
		switch hash {
		case 0:
			var pt entity.Point
			err := json.Unmarshal(*result, &pt)
			utils.ShowError(err, "Unable to Unmarshal Point.")
			pointList = append(pointList, pt)
			break
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
	vertices := entity.InitVertices(pointList)
	return vertices, triangles, normals
}

// getName is a function to get the name of an .obj file.
//
// Parameters:
// 	objPath  - path to the .obj file.
//
// Returns:
//  name - the name of the object.
//
func getName(objPath string) string {
	paths := strings.Split(objPath, "/")
	nameWithExt := paths[len(paths)-1]
	nameWESplitted := strings.Split(nameWithExt, ".")
	name := nameWESplitted[0]
	return name
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
func ReadObj(objPath string) *general.Object {

	// getting abs path
	absPath, err := filepath.Abs(objPath)
	utils.ShowError(err, "Unable to find absolute path for: "+objPath+".")

	// getting the file
	file, err := os.Open(absPath)
	utils.ShowError(err, "Unable to find "+absPath+".")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	vertices, triangles, normals := readLines(scanner)
	name := getName(objPath)

	color := make([]int, 3)
	specularDecay := 100.0
	ambientReflection := 0.0
	diffuseReflection := 0.0
	specularReflection := 0.0
	transReflection := 0.0
	object := general.InitObject(name, vertices, triangles, normals, color, specularDecay, ambientReflection, diffuseReflection, specularReflection, transReflection)

	err = scanner.Err()
	utils.ShowError(err, "Error on reading file: "+absPath+".")
	object.NormalizeNormals()
	object.CheckIntegrity()
	return &object
}
