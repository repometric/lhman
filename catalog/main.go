package catalog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
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
