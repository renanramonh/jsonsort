package file

import (
	"path/filepath"
)

func GetFileNameWithoutExtension(fileName string) string {
	extension := filepath.Ext(fileName)
	return fileName[:len(fileName)-len(extension)]
}
