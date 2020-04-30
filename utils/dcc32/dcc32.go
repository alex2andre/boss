// +build !windows

package dcc32

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashload/boss/msg"
)

func GetDcc32DirByCmd() []string {
	command := exec.Command("where", "dcc32")
	output, err := command.Output()
	if err != nil {
		msg.Warn("dcc32 not found")
	}
	outputStr := strings.ReplaceAll(string(output), "\t", "")
	outputStr = strings.ReplaceAll(outputStr, "\r", "")
	if strings.HasSuffix(outputStr, "\n") {
		outputStr = outputStr[0 : len(outputStr)-1]
	}
	if len(outputStr) == 0 {
		return []string{}
	}
	installations := strings.Split(outputStr, "\n")
	for key, value := range installations {
		installations[key] = filepath.Dir(value)
	}
	return installations
}

func GetDelphiVersionFromRegistry() map[string]string {
	return map[string]string{}
}

func GetDelphiVersionNumberName(currentPath string) string {
	for version, path := range GetDelphiVersionFromRegistry() {
		if strings.HasPrefix(strings.ToLower(path), strings.ToLower(currentPath)) {
			return version
		}
	}
	return ""
}

func GetDelphiPathsByRegistry() []string {
	var paths []string
	for _, path := range GetDelphiVersionFromRegistry() {
		paths = append(paths, filepath.Dir(path))
	}
	return paths
}
