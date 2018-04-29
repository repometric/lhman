package catalog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const catalogFolder = "hub"

// Catalog type represents array of engines
type Catalog []Engine

// Get function creates instance of catalog
func Get() Catalog {
	var result Catalog

	folders, err := ioutil.ReadDir(catalogFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, engineFolder := range folders {

		engine := Engine{}
		depsFilePath := path.Join(catalogFolder, engineFolder.Name(), "deps.json")
		json.Unmarshal(readFile(depsFilePath), &engine.Deps)

		metaFilePath := path.Join(catalogFolder, engineFolder.Name(), "meta.json")
		json.Unmarshal(readFile(metaFilePath), &engine.Meta)

		result = append(result, engine)
	}

	return result
}

//readFile - read file by path, and return []byte
func readFile(path string) []byte {
	result, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("Catch error while reading file: %v\n", e)
		os.Exit(1)
	}
	return result
}

// Recommend function creates instance of catalog with recommendations for installation
func Recommend(projectPath string) Catalog {
	var result Catalog
	extensions := make(map[string]bool)
	var temp[] string
	filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			temp = strings.SplitAfter(path, ".")
			extensions["."+temp[len(temp)-1]] = true
		}
		return nil
	})
	//get list of needable languages
	langs := make(map[string]bool)
	dataLang, e := ioutil.ReadFile(catalogFolder+"/language.linguist.json")
	if e != nil {
		fmt.Printf("Catch error while reading file: %v\n", e)
		os.Exit(1)
	}
	var l Languages
	if json.Unmarshal(dataLang, &l) != nil {
		fmt.Printf("Catch error while parsing file.\n")
		os.Exit(1)
	}

	for _,language:= range l.Definitions.Lang.Language {
		for _, ext := range language.Extensions {
			for extension := range extensions {
				if strings.Contains(ext, extension) {
					for _, name := range language.Name {
						langs[name] = true
					}
				}
			}
		}
	}

	folders, err := ioutil.ReadDir(catalogFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, engineFolder := range folders {
		if engineFolder.IsDir() {
			engine := Engine{}
			depsFilePath := path.Join(catalogFolder, engineFolder.Name(), "deps.json")
			json.Unmarshal(readFile(depsFilePath), &engine.Deps)

			metaFilePath := path.Join(catalogFolder, engineFolder.Name(), "meta.json")
			json.Unmarshal(readFile(metaFilePath), &engine.Meta)

			for extension := range extensions {
				//recommendation by extension
				for _, ext := range engine.Meta.Extensions {
					if strings.Contains(ext, extension) {
						engine.Meta.Recommendation = "Recommended because of the file extension"
					}
				}
				//recommendation by language
				for _, language := range engine.Meta.Languages {
					for lang:=range langs {
						if lang == language {
							engine.Meta.Recommendation = "Recommended because of the code's language"
						}
					}
				}
			}

			result = append(result, engine)
		}
	}
	return result
}