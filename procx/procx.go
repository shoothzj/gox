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
		content, err := os.ReadFile("/proc/" + info.Name() + "/cmdline")
		if err != nil {
			// the proc may end during the loop, ignore it
			continue
		}
		if len(content) == 0 {
			continue
		}
		split := strings.Split(string(bytes.TrimRight(content, string("\x00"))), string(byte(0)))
		cmdLine := strings.Join(split, " ")
		if strings.Contains(cmdLine, processName) {
			result = true
		}
	}
	return result, err
}
