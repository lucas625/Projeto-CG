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
//  SpecularDecay      - constant for how fast the specular component decays.
//  LightPosition      - 3D position of the light.
//  Radius             - radius of the light.
//
type Light struct {
	AmbientIntensity  utils.Vector
	AmbientReflection utils.Vector
	LightIntensity    utils.Vector
	SpecularDecay     float64
	LightPosition     entity.Point
	Radius            float64
}

// Evaluate is a function to evaluate the light at a point.
//
// Parameters:
// 	surfaceNormal - surface normal at the point.
//  camPos        - position of the camera.
//  pos           - the point.
//
// Returns:
// 	the RGB value of the point as a vector.
//
// func (lgt *Light) Evaluate(surfaceNormal utils.Vector, camPos, pos entity.Point) utils.Vector {
// 	// pt to cam vector
// 	ptToCam := entity.ExtractVector(&pos, &camPos)
// 	ptToCamNormalized := utils.NormalizeVector(&ptToCam)
// 	// pt to light
// 	lgtPos := lgt.LightPosition
// 	ptToLight := entity.ExtractVector(&pos, &lgtPos)
// 	ptToLightNormalized := utils.NormalizeVector(&ptToLight)
// 	// R vector
// 	vectorRAux := utils.CMultVector(&surfaceNormal, 2*utils.DotProduct(&surfaceNormal, &ptToLight))
// 	vectorR := utils.SumVector(&vectorRAux, &ptToLightNormalized, 1, -1)
// 	vectorRNormalized := utils.NormalizeVector(&vectorR)
// 	// calculating light
// 	ambientalPart := utils.InitVector(3)
// 	diffusePart := utils.InitVector(3)
// 	specularPart := utils.InitVector(3)
// 	// auxiliars
// 	cosO := utils.DotProduct(&surfaceNormal, &ptToLightNormalized)
// 	cosAWithDecay := math.Pow(utils.DotProduct(&vectorRNormalized, &ptToCamNormalized), lgt.SpecularDecay)
// 	for i := 0; i < 3; i++ {
// 		ambientalPart.Coordinates[i] = lgt.AmbientIntensity.Coordinates[i] * lgt.AmbientReflection.Coordinates[i]
// 		diffusePart.Coordinates[i] = lgt.LightIntensity.Coordinates[i] * lgt.DiffuseReflection.Coordinates[i] * cosO
// 		specularPart.Coordinates[i] = lgt.LightIntensity.Coordinates[i] * lgt.SpecularReflection.Coordinates[i] * cosAWithDecay
// 	}
// 	lightAD := utils.SumVector(&ambientalPart, &diffusePart, 1, 1)
// 	resultingLight := utils.SumVector(&lightAD, &specularPart, 1, 1)
// 	return resultingLight
// }

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
//  radius        - the radius of the light.
//
// Returns:
// 	A Light.
//
func InitLight(lightPos entity.Point, specularDecay, radius float64) Light {
	lgt := Light{
		AmbientIntensity:  utils.InitVector(3),
		AmbientReflection: utils.InitVector(3),
		LightIntensity:    utils.InitVector(3),
		SpecularDecay:     specularDecay,
		Radius:            radius,
		LightPosition:     lightPos}
	return lgt
}
