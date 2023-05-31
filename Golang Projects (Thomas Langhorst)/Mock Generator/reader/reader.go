package reader

import (
	"bufio"
	"os"
	"strings"
)

func Read(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	interfaceLines := make([]string, 0)
	isInterfaceDefiniton := false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if isInterfaceDefinitonStart(line) {
			isInterfaceDefiniton = true
		}

		if isInterfaceDefiniton {
			interfaceLines = append(interfaceLines, line)
		}

		if isInterfaceDefiniton && isInterfaceDefinitonEnd(line) {
			isInterfaceDefiniton = false
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return interfaceLines, nil
}

func isInterfaceDefinitonStart(line string) bool {
	return strings.Contains(line, "type") && strings.Contains(line, "interface") && strings.Contains(line, "{")
}

func isInterfaceDefinitonEnd(line string) bool {
	return line == "}"
}
