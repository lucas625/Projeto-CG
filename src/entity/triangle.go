package entity

// Triangle is a class for a triangle.
//
// Members:
// 	Vertices - list of points indices on a Vertices object.
//
type Triangle struct {
	Vertices []int
}

// InitTriangle is a function to initialize a Triangle.
//
// Parameters:
// 	points - a list of points indices on a Vertices object.
//
// Returns:
// 	a Triangle.
//
func InitTriangle(points []int) Triangle {
	triangle := Triangle{Vertices: points}
	return triangle
}
