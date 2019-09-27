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
//  Camera    - pointer to the camera.
//
type Object struct {
	Name      string
	Vertices  *entity.Vertices
	Triangles *[]entity.Triangle
	Normals   *[]utils.Vector
	Camera    *camera.Camera
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
	utils.ShowError(err, "Unable to get absolute path of object.")
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		utils.ShowError(err, "Unable to create dirs.")
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	utils.ShowError(err, "Unable to get write object.")
}

// InitObject is a function to initialize a Object.
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
func InitObject(name string, vertices *entity.Vertices, triangles *[]entity.Triangle, normals *[]utils.Vector, cam *camera.Camera) Object {
	obj := Object{Name: name, Vertices: vertices, Triangles: triangles, Normals: normals, Camera: cam}
	return obj
}
