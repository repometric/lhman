package catalog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const linterhub = "./linterhub/engine"

// Catalog type represents array of engines
type Catalog []Engine

// Get function creates instance of catalog
func Get() Catalog {
	var result Catalog

	dirs, err := ioutil.ReadDir(linterhub)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dirs {
		if d.IsDir() {
			engineName := d.Name()
			enginePath := filepath.Join(linterhub, engineName)
			metaPath, _ := filepath.Abs(filepath.Join(enginePath, "meta.json"))
			metaFile, e := ioutil.ReadFile(metaPath)
			if e != nil {
				fmt.Printf("Catch error while reading meta file: %v\n", e)
				os.Exit(1)
			}
			depsPath, _ := filepath.Abs(filepath.Join(enginePath, "deps.json"))
			depsFile, e := ioutil.ReadFile(depsPath)
			if e != nil {
				fmt.Printf("Catch error while reading meta file: %v\n", e)
				os.Exit(1)
			}
			var engine Engine
			json.Unmarshal(metaFile, &(engine.Meta))
			json.Unmarshal(depsFile, &(engine.Deps))
			result = append(result, engine)
		}
	}
	return result
}
