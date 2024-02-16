package envfile

import (
	"fmt"
	"os"
	"strings"
)

func ParseEnvFile(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	for ind, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		envar := strings.Split(line, "=")

		if len(envar) != 2 {
			fmt.Printf("Invalid line at line %d\n", ind)
			continue
		}

		key := strings.TrimSpace(envar[0])
		value := strings.TrimSpace(envar[1])

		err := os.Setenv(key, value)
		if err != nil {
			fmt.Printf("Error while setting envar %s, with value %s: %s", key, value, err)
		}
	}
}
