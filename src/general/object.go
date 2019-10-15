package general

import (
	"encoding/json"
	"errors"
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
//  none
//
// Returns:
//  none
//
func (obj *Object) FindCamera(ptCamera *entity.Point) {
	bb := obj.GetBoundingBox()
	vList := make([]float64, 3)
	for i := 0; i < 3; i++ {
		vList[i] = bb[i+3] - bb[i]
	}
	ptTarget := entity.Point{Coordinates: vList}
	cam := camera.InitCameraWithPoints(ptCamera, &ptTarget)
	obj.Camera = &cam
}

// LoadJSONCamera is a function to read all Camera data as json.
//
// Parameters:
//  inPath - path to the input file.
//
// Returns:
//  none
//
func (obj *Object) LoadJSONCamera(inPath string) {
	// opening the file
	camFile, err := os.Open(inPath)
	utils.ShowError(err, "Unable to open camera.")
	// converting to cam
	byteCamera, err := ioutil.ReadAll(camFile)
	utils.ShowError(err, "Unable to convert camera file to bytes.")
	var camAux camera.Camera
	err = json.Unmarshal(byteCamera, &camAux)
	utils.ShowError(err, "Failed to unmarshal camera.")
	// Validating the camera
	if len(camAux.Look.Coordinates) == 0 || len(camAux.Up.Coordinates) == 0 || len(camAux.Right.Coordinates) == 0 {
		if len(camAux.Pos.Coordinates) == 3 {
			obj.FindCamera(&(camAux.Pos))
		} else {
			utils.ShowError(errors.New("Invalid camera"), "Camera with vectors as empty list, but with non 3D position.")
		}
	} else if len(camAux.Look.Coordinates) == 3 && len(camAux.Up.Coordinates) == 3 && len(camAux.Right.Coordinates) == 3 {
		if len(camAux.Pos.Coordinates) == 3 {
			obj.Camera = &camAux
		} else {
			utils.ShowError(errors.New("Invalid camera"), "Camera with vectors with 3D position, but camera position isn't 3D.")
		}
	} else {
		utils.ShowError(errors.New("Invalid camera"), "Camera with invalid vectors.")
	}
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
func InitObject(name string, vertices *entity.Vertices, triangles *[]entity.Triangle, normals *[]utils.Vector, cam *camera.Camera) Object {
	obj := Object{Name: name, Vertices: vertices, Triangles: triangles, Normals: normals, Camera: cam}
	return obj
}
