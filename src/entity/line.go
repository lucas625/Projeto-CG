package entity

import (
	"math"

	"github.com/lucas625/Projeto-CG/src/utils"
)

// Line is a class for lines.
//
// Members:
// 	Start    - the starting point.
// 	Director - the vector director.
//
type Line struct {
	Start    Point
	Director utils.Vector
}

// ExtractLine is a function to extract a line from 2 points.
//
// Parameters:
// 	p0 - the starting point.
//  p1 - the target point.
//
// Returns:
// 	the Line.
//
func ExtractLine(p0, p1 Point) Line {
	v := ExtractVector(&p0, &p1)
	line := Line{Start: p0, Director: v}
	return line
}

// FindPos is a function to get the position of a line at a given t.
//
// Parameters:
// 	t - the t parameter.
//
// Returns:
// 	the Point.
//
func (line Line) FindPos(t float64) Point {
	v := utils.CMultVector(&line.Director, t)
	pos := InitPoint(3)
	for i := 0; i < 3; i++ {
		pos.Coordinates[i] = line.Start.Coordinates[i] + v.Coordinates[i]
	}
	return pos
}

// IntersectPlane is a function to intersect planes.
//
// Parameters:
// 	plane - the Plane.
//
// Returns:
//  The line t parameter (A + tV).
//  A flag checking if has intersection.
//  A flag checking if the plane contains the line.
//
func (line Line) IntersectPlane(plane Plane) (float64, bool, bool) {
	tMult := (line.Director.Coordinates[0] * plane.A) + (line.Director.Coordinates[1] * plane.B) + (line.Director.Coordinates[2] * plane.C)
	cVal := (line.Start.Coordinates[0] * plane.A) + (line.Start.Coordinates[1] * plane.B) + (line.Start.Coordinates[2] * plane.C) + plane.D
	if utils.CheckTolerance(tMult, 0) && utils.CheckTolerance(cVal, 0) {
		return 0, true, true
	} else if utils.CheckTolerance(tMult, 0) && !utils.CheckTolerance(cVal, 0) {
		return 0, false, false
	}
	t := (-1 * cVal) / tMult
	return t, true, false
}

// IntersectSphere is a function to intersect spheres.
//
// Parameters:
// 	plane - the Plane.
//
// Returns:
//  A list with up to 2 line t parameters (A + tV).
//  A flag checking if has intersection.
//
func (line Line) IntersectSphere(sphere Sphere) ([]float64, bool) {
	values := make([]float64, 0, 2)

	i := line.Start.Coordinates[0] - sphere.Center.Coordinates[0]
	j := line.Start.Coordinates[1] - sphere.Center.Coordinates[1]
	k := line.Start.Coordinates[2] - sphere.Center.Coordinates[2]

	t0 := math.Pow(i, 2) + math.Pow(j, 2) + math.Pow(k, 2) - math.Pow(sphere.Radius, 2)
	t1 := 2 * ((line.Director.Coordinates[0] * i) + (line.Director.Coordinates[1] * j) + (line.Director.Coordinates[2] * k))
	t2 := math.Pow(line.Director.Coordinates[0], 2) + math.Pow(line.Director.Coordinates[1], 2) + math.Pow(line.Director.Coordinates[2], 2)

	delta := math.Pow(t1, 2) - (4 * t2 * t0)
	if delta < 0 {
		return values, false
	}
	v1 := ((t1 * -1) + math.Pow(delta, 0.5)) / (2 * t2)
	v2 := ((t1 * -1) - math.Pow(delta, 0.5)) / (2 * t2)

	values = append(values, v1)
	if v1 != v2 {
		values = append(values, v2)
	}
	return values, true
}

// FindBaricentricCoordinates is a function to find the baricentric coordinates of a point.
//
// Parameters:
// 	triang - a list of 3 points.
//  pos    - the target point.
//
// Returns:
//  The 3 baricentric coordinates.
//
func FindBaricentricCoordinates(triang []Point, pos Point) []float64 {

	AB := ExtractVector(&triang[0], &triang[1])
	AC := ExtractVector(&triang[0], &triang[2])

	PA := ExtractVector(&pos, &triang[0])
	PB := ExtractVector(&pos, &triang[1])
	PC := ExtractVector(&pos, &triang[2])

	normal := utils.VectorCrossProduct(&AB, &AC)
	AreaABC := utils.VectorNorm(&normal) / 2

	normalA := utils.VectorCrossProduct(&PB, &PC)
	normalB := utils.VectorCrossProduct(&PC, &PA)

	AreaA := utils.VectorNorm(&normalA) / 2
	AreaB := utils.VectorNorm(&normalB) / 2

	alpha := AreaA / AreaABC
	beta := AreaB / AreaABC
	gama := (1 - alpha) - beta

	coordinates := []float64{alpha, beta, gama}
	return coordinates

}

// IntersectTriangle is a function to intersect triangles.
//
// Parameters:
// 	triang - a list of 3 points.
//
// Returns:
//  The line t parameter (A + tV).
//  The baricentric coordinates at that point.
//  A flag checking if has intersection.
//
func (line Line) IntersectTriangle(triang []Point) (float64, []float64, bool) {
	rayOrigin := line.Start
	rayVector := line.Director

	EPSILON := 0.0000001;

	vertex0 := triang[0]
	vertex1 := triang[1]
	vertex2 := triang[2]


	bCoord := make([]float64, 3)


	edge1 := ExtractVector(&vertex0, &vertex1)
	edge2 := ExtractVector(&vertex0, &vertex2)

    h := utils.VectorCrossProduct(&rayVector, &edge2)
	a := utils.DotProduct(&edge1, &h);
	
    if a > -EPSILON && a < EPSILON {
		return 0,bCoord, false;    // This ray is parallel to this triangle.
	}
	f := 1.0/a
	
    s := ExtractVector(&vertex0,&rayOrigin)
    u := f * utils.DotProduct(&s, &h)
    if u < 0.0 || u > 1.0 {
		return 0,bCoord,false
	}
    q := utils.VectorCrossProduct(&s,&edge1);
    v := f * utils.DotProduct(&rayVector,&q);
    if v < 0.0 || u + v > 1.0 {
		return 0,bCoord,false
	}
    // At this stage we can compute t to find out where the intersection point is on the line.
    t := f * utils.DotProduct(&edge2, &q)
    if t > EPSILON && t < 1/EPSILON {
		// outIntersectionPoint = rayOrigin + rayVector * t
		bCoord = []float64{1-u-v, u, v}
        return t,bCoord,true
    }
    
	return 0,bCoord,false
}
