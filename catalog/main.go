package catalog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const bundlePath = "tmp/bundle.json"

// Catalog type represents array of engines
type Catalog []Engine

// Get function creates instance of catalog
func Get() Catalog {
	var result Catalog

	bundleFile, e := ioutil.ReadFile(bundlePath)
	if e != nil {
		fmt.Printf("Catch error while reading bundle file: %v\n", e)
		os.Exit(1)
	}

	type byteArray []byte

	var m = make(map[string]interface{})
	if json.Unmarshal(bundleFile, &m) != nil {
		fmt.Printf("Catch error while parsing bundle file.\n")
		os.Exit(1)
	}

	for key, value := range m {
		if !strings.Contains(key, "$") {
			strB, _ := json.Marshal(value)
			var engine Engine
			json.Unmarshal(strB, &(engine))
			result = append(result, engine)
		}
	}

	return result
}
