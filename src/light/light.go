package light

import (
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Light is a structure for holding all light data.
//
// Members:
//  AmbientIntensity   - RGB for the ambient intensity.
//  AmbientReflection  - RGB for the ambient reflection.
//  LightIntensity     - RGB for the light intensity.
//  DiffuseReflection  - RGB for how diffuse is the object.
//  SpecularReflection - RGB for how specular is the object.
//
type Light struct {
	AmbientIntensity   utils.Vector
	AmbientReflection  utils.Vector
	LightIntensity     utils.Vector
	DiffuseReflection  utils.Vector
	SpecularReflection utils.Vector
}
