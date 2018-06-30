package structs

import "strings"

type ConfigParser interface {
	ParseConfig([]byte) interface{}
}

func GetParser(filePath string) ConfigParser {
	for dir, parser := range StructToConfig {
		if !strings.HasPrefix(filePath, dir) {
			continue
		}

		return parser
	}

	return nil
}
