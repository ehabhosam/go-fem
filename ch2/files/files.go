package files

import (
	"os"
)

func ReadTextFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		// something went wrong
		return "", err
	} else {
		// kolo tmam
		return string(content), nil
	}
}

func WriteToTextFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644) // already returns an error if something goes wrong
}
