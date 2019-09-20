package obj

import (
	"io/ioutil"
	"path/filepath"
)

func readObj(objPath string) {
	absPath, err := filepath.Abs(objPath)
	utils.showError(err, "Unable to find absolute path for "+objPath)
	data, err := ioutil.ReadFile(absPath)
}
