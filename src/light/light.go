package light

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/lucas625/Projeto-CG/src/entity"
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
//  SpecularDecay      - constant for how fast the specular component decays.
//  LightPosition      - 3D position of the light.
//
type Light struct {
	AmbientIntensity   utils.Vector
	AmbientReflection  utils.Vector
	LightIntensity     utils.Vector
	DiffuseReflection  utils.Vector
	SpecularReflection utils.Vector
	SpecularDecay      float64
	LightPosition      entity.Point
}

// Avaliate is a function to avaliate the light at a point.
//
// Parameters:
// 	surfaceNormal - surface normal at the point.
//  pointToCam    - vector from the point to the camera.
//  pos           - the point.
//
// Returns:
// 	the RGB value of the point
//
func (lgt *Light) Avaliate(surfaceNormal, pointToCam utils.Vector, pos entity.Point) {

}

// LoadJSONLight is a function to read all Light data as json.
//
// Parameters:
//  inPath - path to the input file.
//
// Returns:
//  the light
//
func LoadJSONLight(inPath string) *Light {
	// opening the file
	lightFile, err := os.Open(inPath)
	utils.ShowError(err, "Unable to open light.")
	// converting to light
	byteLight, err := ioutil.ReadAll(lightFile)
	utils.ShowError(err, "Unable to convert light file to bytes.")
	var lightAux Light
	err = json.Unmarshal(byteLight, &lightAux)
	utils.ShowError(err, "Failed to unmarshal light.")
	// Validating the camera
	if len(lightAux.AmbientIntensity.Coordinates) != 3 ||
		len(lightAux.AmbientReflection.Coordinates) != 3 ||
		len(lightAux.LightIntensity.Coordinates) != 3 ||
		len(lightAux.DiffuseReflection.Coordinates) != 3 ||
		len(lightAux.SpecularReflection.Coordinates) != 3 ||
		len(lightAux.LightPosition.Coordinates) != 3 {
		utils.ShowError(errors.New("invalid length of a light component"), "light components must have length equal to 3.")
	}
	return &lightAux
}

// WriteJSONLight is a function to write all Light data as json.
//
// Parameters:
//  outPath - path to the output folder.
//
// Returns:
//  none
//
func (lgt *Light) WriteJSONLight(outPath string) {
	// creating the json
	file, err := json.MarshalIndent(*lgt, "", "	")
	utils.ShowError(err, "Unable to convert light to json.")
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, "light.json"))
	utils.ShowError(err, "Unable to get light's absolute path.")
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		utils.ShowError(err, "Unable to create dirs.")
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	utils.ShowError(err, "Unable to write light.")
}

// InitLight is a function to initialize a Light.
//
// Parameters:
// 	lightPos      - the position of the light.
// 	specularDecay - how fast the specular component decays.
//
// Returns:
// 	A Light.
//
func InitLight(lightPos entity.Point, specularDecay float64) Light {
	lgt := Light{
		AmbientIntensity:   utils.InitVector(3),
		AmbientReflection:  utils.InitVector(3),
		LightIntensity:     utils.InitVector(3),
		DiffuseReflection:  utils.InitVector(3),
		SpecularReflection: utils.InitVector(3),
		SpecularDecay:      specularDecay,
		LightPosition:      lightPos}
	return lgt
}
