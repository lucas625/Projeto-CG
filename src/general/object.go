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

// Objects is a class for all objects.
//
// Members:
// Label   - the label.
// ObjList - a list of objects.
//
type Objects struct {
	Label   string
	ObjList []Object
}

// WriteJSONObjects is a function to write all Object data as json.
//
// Parameters:
//  outPath - path to the output folder.
//
// Returns:
//  none
//
func (objs *Objects) WriteJSONObjects(outPath string) {
	// creating the json
	file, err := json.MarshalIndent(*objs, "", "	")
	utils.ShowError(err, "Unable to convert object to json.")
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, objs.Label+".json"))
	utils.ShowError(err, "Unable to get objects's absolute path.")
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		utils.ShowError(err, "Unable to create dirs.")
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	utils.ShowError(err, "Unable to write objects.")
}

// LoadJSONObjects is a function to read all Objects data as json.
//
// Parameters:
//  inPath - path to the input file.
//
// Returns:
//  the objects.
//
func LoadJSONObjects(inPath string) *Objects {
	// opening the file
	objsFile, err := os.Open(inPath)
	utils.ShowError(err, "Unable to open objects.")
	// converting to cam
	byteCamera, err := ioutil.ReadAll(objsFile)
	utils.ShowError(err, "Unable to convert objects file to bytes.")
	var objsAux Objects
	err = json.Unmarshal(byteCamera, &objsAux)
	utils.ShowError(err, "Failed to unmarshal objects.")
	for _, obj := range objsAux.ObjList {
		obj.CheckIntegrity()
	}
	return &objsAux
}

// InitObjects is a function to initialize a objects.
//
// Parameters:
//  none
//
// Returns:
//  a objects.
//
func InitObjects(label string, objlist []Object) *Objects {
	objs := Objects{Label: label, ObjList: objlist}
	return &objs
}

// Object is a class for all data of an object.
//
// Members:
//  Name               - the name of the object.
// 	Vertices           - pointer to a Vertices object.
// 	Triangles          - pointer to a list with all Triangles of the object.
//  Normals            - pointer to a list of Vectors with all vertices normals.
//  Color              - RGB for the color of the object.
//  SpecularDecay      - constant for how fast the specular component decays.
//  SpecularReflection - the coeficient of specular reflection.
//  TransReflection    - the coeficient for transmission.
//  AmbientReflection  - RGB for the ambient reflection.
//  DiffuseReflection  - Diffuse reflection coeficient.
//
type Object struct {
	Name               string
	Vertices           entity.Vertices
	Triangles          []entity.Triangle
	Normals            []utils.Vector
	Color              []int
	SpecularDecay      float64
	SpecularReflection float64
	TransReflection    float64
	AmbientReflection  float64
	DiffuseReflection  float64
}

// CheckIntegrity is a function to check the attributes of an object.
//
// Parameters:
// 	none
//
// Returns:
//  none
//
func (obj *Object) CheckIntegrity() {
	for _, vertice := range obj.Vertices.Points {
		if len(vertice.Coordinates) != 3 {
			utils.ShowError(errors.New("Invalid object"), "Vertices length not equal 3.")
		}
	}
	for _, triangle := range obj.Triangles {
		if len(triangle.Vertices) != 3 {
			utils.ShowError(errors.New("Invalid object"), "Triangle with number of vertices not equal 3.")
		}
		if len(triangle.Normals) != 3 {
			utils.ShowError(errors.New("Invalid object"), "Triangle with number of normals not equal 3.")
		}
	}
	for _, normal := range obj.Normals {
		if len(normal.Coordinates) != 3 {
			utils.ShowError(errors.New("Invalid object"), "Normal length not equal 3.")
		}
	}
	if len(obj.Color) != 3 {
		utils.ShowError(errors.New("Invalid object"), "Color length not equal 3.")
	}
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
	normals := obj.Normals
	for i := range obj.Normals {
		auxV := utils.NormalizeVector(&normals[i])
		obj.Normals[i] = auxV
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
//  name               - the name of the object.
// 	vertices           - Vertices object with all points.
//  triangles          - list of triangles.
//  normals            - list of normal vectors.
//  color              - the color of the object.
//  specularDecay      - constant for how fast the specular component decays.
//  specularReflection - the coeficient of specular reflection.
//  TransReflection    - the coeficient for transmission.
//
// Returns:
//  the object.
//
func InitObject(name string, vertices entity.Vertices, triangles []entity.Triangle, normals []utils.Vector, color []int, specularDecay, ambientReflection, diffuseReflection, specularReflection, transReflection float64) Object {
	obj := Object{Name: name, Vertices: vertices, Triangles: triangles, Normals: normals, Color: color, SpecularDecay: specularDecay, AmbientReflection: ambientReflection, DiffuseReflection: diffuseReflection, SpecularReflection: specularReflection, TransReflection: transReflection}
	obj.CheckIntegrity()
	return obj
}
