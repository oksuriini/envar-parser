package envfile

import (
	"fmt"
	"os"
	"strings"
)

// Parses environmental variables from file and sets them as envars
func SetAllEnvars(fileName string) {
	lines, err := parseFile(fileName)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for index, line := range lines {
		key, value := parseLine(line, index)
		err := os.Setenv(key, value)
		if err != nil {
			fmt.Printf("Error while setting envar %s, with value %s: %s", key, value, err)
		}
	}
}

// Parses environmental variables from file and returns the value of found key
// If no key is found then returns empty string
func GetOneEnvar(envName, fileName string) (string, error) {
	lines, err := parseFile(fileName)
	if err != nil {
		fmt.Println("error: ", err)
		return "", err
	}

	for index, line := range lines {
		key, value := parseLine(line, index)
		if key == envName {
			return value, nil
		} else {
			continue
		}
	}

	err = fmt.Errorf("could not find envar %s", envName)
	return "", err
}

// Parses given line and returns found key and value
func parseLine(line string, index int) (key, value string) {
	if line == "" || strings.HasPrefix(line, "#") {
		return "", ""
	}

	envar := strings.Split(line, "=")

	if len(envar) != 2 {
		fmt.Printf("Invalid line at line %d\n", index)
		return "", ""
	}

	key = strings.TrimSpace(envar[0])
	value = strings.TrimSpace(envar[1])
	return
}

// Parses file and returns slice of strings that hold all parsed lines
func parseFile(fileName string) ([]string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}
