package entity

// Triangle is a class for a triangle.
//
// Members:
// 	Vertices - list of points indices on a Vertices object.
//  Normals  - list of normal vectors indices.
//
type Triangle struct {
	Vertices []int
	Normals  []int
}

// InitTriangle is a function to initialize a Triangle.
//
// Parameters:
// 	points  - a list of points indices on a Vertices object.
//  normals - a list of normal vectors indices.
//
// Returns:
// 	a Triangle.
//
func InitTriangle(points, normals []int) Triangle {
	triangle := Triangle{Vertices: points, Normals: normals}
	return triangle
}
