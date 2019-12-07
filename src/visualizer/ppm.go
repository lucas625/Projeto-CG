package visualizer

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/lucas625/Projeto-CG/src/screen"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// WritePPM is a function to write a ppm.
//
// Parameters:
//  sc      - the colored screen.
//  outPath - path to the output folder.
//
// Returns:
//  none
//
func WritePPM(sc screen.ColoredScreen, outPath string) {
	ppmAsString := ""
	header := "P3\n# object.ppm\n" + strconv.Itoa(sc.Width) + " " + strconv.Itoa(sc.Height) + "\n255\n"
	body := ""
	for i := 0; i < sc.Height; i++ {
		for j := 0; j < sc.Width; j++ {
			for k := 0; k < 3; k++ {
				body = body + strconv.Itoa(sc.Colors[i][j][k])
				if k+1 < 3 {
					body = body + "  "
				} else if j+1 < sc.Width {
					body = body + "    "
				} else {
					body = body + "\n"
				}
			}
		}
	}
	ppmAsString = header + body
	Write(outPath, ppmAsString)
}

// Write is a function to write a ppm with its string.
//
// Parameters:
//  outPath     - path to the output folder.
//  ppmAsString - the ppm formated as string.
//
// Returns:
//  none
//
func Write(outPath, ppmAsString string) {
	// creating the json
	file := []byte(ppmAsString)
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, "object.ppm"))
	utils.ShowError(err, "Unable to get objects's absolute path.")
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		utils.ShowError(err, "Unable to create dirs.")
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	utils.ShowError(err, "Unable to write objects.")
}
