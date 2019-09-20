package general

import (
	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Object is a class for all data of an object.
//
// Members:
// 	vertices  - pointer to a Vertices object.
// 	triangles - pointer to a list with all Triangles of the object.
//  normals   - pointer to a list of Vectors with all vertices normals.
//  camera    - pointer to the camera.
//
type Object struct {
	Vertices  *entity.Vertices
	Triangles *[]entity.Triangle
	Normals   *[]utils.Vector
	Camera    *camera.Camera
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
func InitObject(vertices *entity.Vertices, triangles *[]entity.Triangle, normals *[]utils.Vector, cam *camera.Camera) Object {
	obj := Object{Vertices: vertices, Triangles: triangles, Normals: normals, Camera: cam}
	return obj
}
