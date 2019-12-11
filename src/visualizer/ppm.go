package visualizer

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"fmt"

	"github.com/lucas625/Projeto-CG/src/screen"
	"github.com/lucas625/Projeto-CG/src/utils"
)

// WritePPM is a function to write a ppm.
//
// Parameters:
//  sc      - the colored screen.
//  outPath - path to the output folder.
//  name    - the name of the obj.
//  p       - flag to print (if true will not write).
//
// Returns:
//  none
//
func WritePPM(sc screen.ColoredScreen, outPath, name string, p bool) {
	ppmAsString := ""
	header := "P3\n# "+name+".ppm\n" + strconv.Itoa(sc.Width) + " " + strconv.Itoa(sc.Height) + "\n255\n"
	body := ""
	count := 0
	for i := 0; i < sc.Height; i++ {
		body = body + " "
		for j := 0; j < sc.Width; j++ {
			for k := 0; k < 3; k++ {
				body = body + strconv.Itoa(sc.Colors[i][j][k])
				if k+1 < 3 {
					body = body + "  "
				} else if count < 3 {
					body = body + "    "
				}
			}
			count++
			if count >= 4 {
				body = body + "\n"
				count = 0
				if j+1 < sc.Width {
					body = body + " "
				}
			}
		}
	}
	if count != 0 {
		body = body + "\n"
		count = 0
	}
	ppmAsString = header + body
	if p {
		fmt.Println(ppmAsString)
	}else {
		Write(outPath, ppmAsString, name)
	}
	
}

// Write is a function to write a ppm with its string.
//
// Parameters:
//  outPath     - path to the output folder.
//  ppmAsString - the ppm formated as string.
//  name        - the name of the file.
//
// Returns:
//  none
//
func Write(outPath, ppmAsString, name string) {
	// creating the json
	file := []byte(ppmAsString)
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, name+".ppm"))
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
