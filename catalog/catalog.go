package catalog

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
	"encoding/json"
)

const linterhub = "./linterhub/engine"

type Catalog []Engine

func Get() Catalog {
	var result Catalog

	dirs, err := ioutil.ReadDir(linterhub)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dirs {
		if d.IsDir() {
			engineName := d.Name()
			enginePath := filepath.Join(linterhub,engineName)
			metaPath, _ := filepath.Abs(filepath.Join(enginePath, "meta.json"))
			metaFile, e := ioutil.ReadFile(metaPath)
			if e != nil {
				fmt.Printf("Catch error while reading meta file: %v\n", e)
				os.Exit(1)
			}
			var engine Engine
			json.Unmarshal(metaFile, &(engine.Meta))
			result = append(result, engine)
		}
	}
	return result
}

