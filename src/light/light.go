package light

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/lucas625/Projeto-CG/src/general"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Lights is a structure for holding all light data.
//
// Members:
//  LightList - the list of lights.
//
type Lights struct {
	LightList []Light
}

// LoadJSONLights is a function to read all Light data as json.
//
// Parameters:
//  inPath - path to the input file.
//
// Returns:
//  the light
//
func LoadJSONLights(inPath string) *Lights {
	// opening the file
	lightFile, err := os.Open(inPath)
	utils.ShowError(err, "Unable to open light.")
	// converting to light
	byteLight, err := ioutil.ReadAll(lightFile)
	utils.ShowError(err, "Unable to convert light file to bytes.")
	var lightAux Lights
	err = json.Unmarshal(byteLight, &lightAux)
	utils.ShowError(err, "Failed to unmarshal light.")
	// Validating the camera
	for _, lgt := range lightAux.LightList {
		if len(lgt.Color) != 3 {
			utils.ShowError(errors.New("invalid length of a light component"), "light components must have length equal to 3.")
		}
	}

	return &lightAux
}

// WriteJSONLights is a function to write all Light data as json.
//
// Parameters:
//  outPath - path to the output folder.
//
// Returns:
//  none
//
func (lgt *Lights) WriteJSONLights(outPath string) {
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

// InitLights is a function to initialize a Lights.
//
// Parameters:
// 	lightList - the list of lights.
//
// Returns:
// 	the Lights object.
//
func InitLights(lightList []Light) Lights {
	lights := Lights{LightList: lightList}
	return lights
}

// Light is a structure for holding a light data.
//
// Members:
//  AmbientIntensity   - the ambient intensity.
//  LightIntensity     - the light intensity.
//  LightObject        - Object for the light.
//  Color              - RGB of the light.
//
type Light struct {
	AmbientIntensity float64
	LightIntensity   float64
	Color            []int
	LightObject      general.Object
}

// InitLight is a function to initialize a Light.
//
// Parameters:
//  ambientIntensity  - the ambient intensity of the light.
//  lightIntensity    - the instensity of the light.
//  object            - the object tha defines the light.
//  color             - the rgb of the light.
//
// Returns:
// 	A Light.
//
func InitLight(ambientIntensity, lightIntensity float64, object general.Object, color []int) Light {
	lgt := Light{
		AmbientIntensity: ambientIntensity,
		LightIntensity:   lightIntensity,
		LightObject:      object,
	}
	return lgt
}
