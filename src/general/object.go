package general

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Object is a class for all data of an object.
//
// Members:
//  Name      - the name of the object.
// 	Vertices  - pointer to a Vertices object.
// 	Triangles - pointer to a list with all Triangles of the object.
//  Normals   - pointer to a list of Vectors with all vertices normals.
//  DiffuseReflection  - RGB for how diffuse is the object.
//  SpecularReflection - RGB for how specular is the object.
//
type Object struct {
	Name               string
	Vertices           *entity.Vertices
	Triangles          *[]entity.Triangle
	Normals            *[]utils.Vector
	DiffuseReflection  utils.Vector
	SpecularReflection utils.Vector
}

// WriteJSONObject is a function to write all Object data as json.
//
// Parameters:
// 	obj      - a list of points.
//  outPath - path to the output folder.
//
// Returns:
//  none
//
func WriteJSONObject(obj *Object, outPath string) {
	// creating the json
	file, err := json.MarshalIndent(*obj, "", "	")
	utils.ShowError(err, "Unable to convert object to json.")
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, obj.Name+".json"))
	utils.ShowError(err, "Unable to get object's absolute path.")
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		utils.ShowError(err, "Unable to create dirs.")
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	utils.ShowError(err, "Unable to write object.")
}

// NormalizeNormals is a function to normalize all triangle normals.
//
// Parameters:
// 	none
//
// Returns:
//  none
//
func (obj *Object) NormalizeNormals() {
	normals := *obj.Normals
	for i := range *obj.Normals {
		auxV := utils.NormalizeVector(&normals[i])
		normals[i] = auxV
	}
}

// GetBoundingBox is a function to get the bounding box of an Object.
//
// Parameters:
//  none
//
// Returns:
//  [minX, minY, minZ, maxX, maxY, maxZ]
//
func (obj *Object) GetBoundingBox() []float64 {
	vertices := obj.Vertices
	bb := make([]float64, 6)
	for i := 0; i < 3; i++ {
		bb[i] = vertices.Points[0].Coordinates[i]
		bb[i+3] = vertices.Points[0].Coordinates[i]
	}
	for _, point := range vertices.Points {
		for j := 0; j < 3; j++ {
			if bb[j] > point.Coordinates[j] {
				bb[j] = point.Coordinates[j]
			}
			if bb[j+3] < point.Coordinates[j] {
				bb[j+3] = point.Coordinates[j]
			}
		}
	}
	return bb
}

// FindCamera is a function to initialize a Camera.
//
// Parameters:
//  ptCamera - the position of the camera.
//
// Returns:
//  the camera.
//
func (obj *Object) FindCamera(ptCamera *entity.Point) *camera.Camera {
	bb := obj.GetBoundingBox()
	vList := make([]float64, 3)
	for i := 0; i < 3; i++ {
		vList[i] = bb[i+3] - bb[i]
	}
	ptTarget := entity.Point{Coordinates: vList}
	cam := camera.InitCameraWithPoints(ptCamera, &ptTarget)
	return &cam
}

// InitObject is a function to initialize an Object.
//
// Parameters:
// 	points - a list of points.
//
// Returns:
//  vertices  - Vertices object with all points.
//  triangles - list of triangles.
//  normals   - list of normal vectors.
//  cam       - the cam to the object.
//
func InitObject(name string, vertices *entity.Vertices, triangles *[]entity.Triangle, normals *[]utils.Vector, diffuseReflection, specularReflection utils.Vector) Object {
	obj := Object{Name: name, Vertices: vertices, Triangles: triangles, Normals: normals, DiffuseReflection: diffuseReflection, SpecularReflection: specularReflection}
	return obj
}
