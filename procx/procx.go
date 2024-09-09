package procx

import (
	"bytes"
	"os"
	"regexp"
	"strings"
)

var pidPattern = regexp.MustCompile("[0-9]+")

func ProcessExists(processName string) (bool, error) {
	result := false
	fileInfos, err := os.ReadDir("/proc")
	if err != nil {
		return false, err
	}
	for _, info := range fileInfos {
		name := info.Name()
		if !pidPattern.MatchString(name) {
			continue
		}
		cmdLine, err := parseCmdLine("/proc/" + info.Name() + "/cmdline")
		if err != nil {
			// the proc may end during the loop, ignore it
			continue
		}
		if strings.Contains(cmdLine, processName) {
			result = true
		}
	}
	return result, err
}

func parseCmdLine(path string) (string, error) {
	cmdData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	if len(cmdData) < 1 {
		return "", nil
	}

	split := strings.Split(string(bytes.TrimRight(cmdData, string("\x00"))), string(byte(0)))
	return strings.Join(split, " "), nil
}
