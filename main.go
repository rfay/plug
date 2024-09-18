package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	// Get the Go module cache path
	goModCache, err := getGoModCache()
	if err != nil {
		panic(err)
	}

	// Construct the GoPath to include both the module cache and the current directory
	goPath := fmt.Sprintf("%s%c%s", goModCache, os.PathListSeparator, ".")

	i := interp.New(interp.Options{
		GoPath:       goPath,  // Set the GoPath to include module cache and current directory
		GoModulePath: ".",     // Set the GoModulePath to the current directory
	})

	// Use the standard library symbols
	i.Use(stdlib.Symbols)

	// Add the symbols from the yaml package
	i.Use(map[string]map[string]reflect.Value{
		"gopkg.in/yaml.v3": {
			"Marshal":    reflect.ValueOf(yaml.Marshal),
			"Unmarshal":  reflect.ValueOf(yaml.Unmarshal),
			"NewDecoder": reflect.ValueOf(yaml.NewDecoder),
			"NewEncoder": reflect.ValueOf(yaml.NewEncoder),
			// Add other functions or types if needed
		},
	})

	// The code to be interpreted
	src := `
		package main

		import (
			"fmt"
			"io/ioutil"
			yaml "gopkg.in/yaml.v3"
		)

		func main() {
			data, err := ioutil.ReadFile("junk.yaml")
			if err != nil {
				panic(err)
			}

			var m map[string]interface{}
			err = yaml.Unmarshal(data, &m)
			if err != nil {
				panic(err)
			}

			// Traverse the map
			for k, v := range m {
				fmt.Printf("%s: %v\n", k, v)
			}
		}
	`

	// Evaluate the code
	_, err = i.Eval(src)
	if err != nil {
		panic(err)
	}

	// Call the main function in the interpreted code
	_, err = i.Eval("main()")
	if err != nil {
		panic(err)
	}
}

func getGoModCache() (string, error) {
	cmd := exec.Command("go", "env", "GOMODCACHE")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
