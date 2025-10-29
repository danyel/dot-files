package configuration

import (
	"os"
	"path/filepath"
)

type DotFileConfiguration struct {
	Project        string
	ConfigFileName string
	GetHomeDirFunc func() (string, error)
}

func ConstructPath(dotFileConfiguration DotFileConfiguration) string {
	GetHomeDirFunc := dotFileConfiguration.GetHomeDirFunc
	if GetHomeDirFunc == nil {
		GetHomeDirFunc = os.UserHomeDir
	}
	dir, err := GetHomeDirFunc()
	if err != nil {
		panic(err)
	}

	return filepath.Join(dir, ".config", dotFileConfiguration.Project, dotFileConfiguration.ConfigFileName)
}
