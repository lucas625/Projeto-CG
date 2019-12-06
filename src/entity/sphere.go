package entity

// Sphere is a class for spheres.
//
// Members:
// 	Center - the center of the sphere.
//  Radius - the radius of the sphere.
//
type Sphere struct {
	Center Point
	Radius float64
}

// InitSphere is a function to initialize a sphere.
//
// Parameters:
// 	center - the center of the sphere.
//  radius - the radius of the sphere.
//
// Returns:
// 	the sphere.
//
func InitSphere(center Point, radius float64) Sphere {
	return Sphere{Center: center, Radius: radius}
}
