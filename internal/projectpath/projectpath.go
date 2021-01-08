package projectpath

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var Root = getRoot()

func getRoot() string {
	if os.Getenv("ENVIRONMENT") == "local" || os.Getenv("ENVIRONMENT") == "debug" || os.Getenv("ENVIRONMENT") == "" {
		_, b, _, _ := runtime.Caller(0)
		return path.Join(path.Dir(b), "../..")
	} else {
		b, _ := os.Getwd()
		return b
	}
}

// project.Path("internal/handler") - returns the absolute path for the given folder
func Path(pathTo string) string {
	return filepath.Join(Root, pathTo)
}
