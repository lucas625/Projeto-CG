package camera

import (
	"fmt"
	"log"

	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Camera is a class for cameras.
//
// Members:
// 	Pos   - the position of the camera
// 	Look  - vector to were the camera is looking.
//  Up    - vector head of the camera.
//  Right - side vector of the camera.
//
type Camera struct {
	Pos   entity.Point
	Look  utils.Vector
	Up    utils.Vector
	Right utils.Vector
}

// CamToHomogeneousMatrix is a function to create the matrix ready(after transposition) to multiply the points.
//
// Parameters:
// 	cam - a Camera.
//
// Returns:
// 	a Matrix.
//
func CamToHomogeneousMatrix(cam *Camera) utils.Matrix {
	maux := utils.InitMatrix(3, 3)
	// placing vectors on the matrix on the right form
	maux.Values[0] = cam.Right.Coordinates
	maux.Values[1] = cam.Up.Coordinates
	maux.Values[2] = cam.Look.Coordinates
	// adding homogeneous and translation
	maux.Values = append(maux.Values, []float64{0, 0, 0, 1})
	pValues := cam.Pos.Coordinates
	for i := 0; i < 3; i++ {
		maux.Values[i] = append(maux.Values[i], pValues[i]*-1)
	}
	maux.Lines++
	maux.Columns++
	return maux
}

// CheckLenVector is a function to check the length of the camera vectors.
//
// Parameters:
// 	a - a vector.
//
// Returns:
// 	none
//
func CheckLenVector(vect utils.Vector) {
	if len(vect.Coordinates) != 3 {
		log.Fatalf("Invalid length of Camera vector %d.\n", len(vect.Coordinates))
	}
}

// InitCamera is a function to initialize a Camera.
//
// Parameters:
// 	pos   - the position of the camera.
// 	look  - vector to were the camera is looking.
//  up    - vector head of the camera.
//  right - side vector of the camera.
//
// Returns:
// 	A Camera.
//
func InitCamera(pos entity.Point, look, up, right utils.Vector) Camera {
	if len(pos.Coordinates) != 3 {
		log.Fatalf("Invalid length of Camera point %d.\n", len(pos.Coordinates))
	}
	CheckLenVector(look)
	CheckLenVector(up)
	CheckLenVector(right)
	cam := Camera{Pos: pos, Look: look, Up: up, Right: right}
	return cam
}

// InitCameraWithPoints is a function to initialize a Camera based only on its position and the target point.
//
// Parameters:
// 	pos    - the position of the camera.
// 	target - target Point.
//
// Returns:
// 	A Camera.
//
func InitCameraWithPoints(pos, target *entity.Point) Camera {
	look := entity.ExtractVector(pos, target)
	look = utils.NormalizeVector(&look)

	vectTemp := utils.Vector{Coordinates: []float64{0, 1, 0}}
	vectTemp = utils.NormalizeVector(&vectTemp)
	right := utils.VectorCrossProduct(&vectTemp, &look)
	right = utils.NormalizeVector(&right)
	fmt.Println("right", right)
	up := utils.VectorCrossProduct(&look, &right)
	up = utils.NormalizeVector(&up)

	return InitCamera(*pos, look, up, right)
}
