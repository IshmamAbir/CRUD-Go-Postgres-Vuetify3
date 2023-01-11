package util

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// ReadConfig function to read configuration from a file
func ReadConfig(path string) (map[string]string, error) {

	var configMap = make(map[string]string)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		keyVal := strings.Split(line, ":=")
		if len(keyVal) != 2 {
			return nil, errors.New("error reading config file.check key vlaue assigner ':='")
		}
		if keyVal[0] != "" && keyVal[1] != "" {
			configMap[keyVal[0]] = keyVal[1]
		}
	}
	return configMap, nil
}
