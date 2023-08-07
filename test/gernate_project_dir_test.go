package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var jsonData map[string]any

const jsonFileName = "dir.json"

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWorkDIr, _ := os.Getwd()
	stRootDir = stWorkDIr[:strings.LastIndex(stWorkDIr, stSeparator)]
	gnJsonBytes, _ := os.ReadFile(stWorkDIr + stSeparator + jsonFileName)
	err := json.Unmarshal(gnJsonBytes, &jsonData)
	if err != nil {
		panic("Load Json Data Error" + err.Error())
	}
}

func parseMap(mapData map[string]any, parentDir string) {
	for k, v := range mapData {
		switch v.(type) {
		case string:
			path, _ := v.(string)
			if path == "" {
				continue
			}
			if parentDir != "" {
				path = parentDir + stSeparator + path
				if k == "text" {
					parentDir = path
				}
			} else {
				parentDir = path
			}
			createDir(path)
		case []any:
			parseArray(v.([]any), parentDir)
		}
	}
}

func parseArray(jsonData []any, parentDir string) {
	for _, v := range jsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, parentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	err := os.MkdirAll(stRootDir+stSeparator+path, fs.ModePerm)
	if err != nil {
		panic("Create Dir Error:" + err.Error())
	}

}

func TestGenerateDir(t *testing.T) {
	loadJson()
	parseMap(jsonData, "")
}
