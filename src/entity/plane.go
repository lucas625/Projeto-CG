package entity

import (
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Plane is a class for planes using formula AX + BY + CZ + D = 0.
//
// Members:
// 	A - X coeficient.
// 	B - Y coeficient.
//  C - Z coeficient.
//  D - constant.
//
type Plane struct {
	A float64
	B float64
	C float64
	D float64
}

// ExtractPlane is a function to extract a plane from 3 points.
//
// Parameters:
// 	p0 - the first point.
//  p1 - the second point.
//  p2 - the third point.
//
// Returns:
// 	the plane.
//
func ExtractPlane(p0, p1, p2 Point) Plane {
	v1 := ExtractVector(&p0, &p1)
	v2 := ExtractVector(&p0, &p2)

	normV := utils.VectorCrossProduct(&v1, &v2)
	a := normV.Coordinates[0]
	b := normV.Coordinates[1]
	c := normV.Coordinates[2]
	d := -1 * ((a * p0.Coordinates[0]) + (b * p0.Coordinates[1]) + (c * p0.Coordinates[2]))

	plane := Plane{A: a, B: b, C: c, D: d}
	return plane

}
