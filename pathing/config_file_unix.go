// +build darwin freebsd linux netbsd openbsd solaris

package pathing

import (
	"os"
	"path/filepath"
)

const (
	defaultConfigFile = "packer"
)

func getDefaultConfigDir() string {

	var defaultConfigFileDir string

	if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
		defaultConfigFileDir = xdgConfigHome
	} else {
		defaultConfigFileDir = filepath.Join(os.Getenv("HOME"), "config")
	}

	return filepath.Join(defaultConfigFileDir, "packer")
}
