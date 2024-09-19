package main

import (
	"os"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	i := interp.New(interp.Options{})

	// Use the standard library symbols
	i.Use(stdlib.Symbols)

	// Register the symbols from the yaml package
	i.Use(map[string]map[string]reflect.Value{
		"gopkg.in/yaml.v3": {
			"Marshal":    reflect.ValueOf(yaml.Marshal),
			"Unmarshal":  reflect.ValueOf(yaml.Unmarshal),
			"NewDecoder": reflect.ValueOf(yaml.NewDecoder),
			"NewEncoder": reflect.ValueOf(yaml.NewEncoder),
			// Add other functions or types if needed
		},
	})

	// Register the os package symbols if needed
	i.Use(map[string]map[string]reflect.Value{
		"os": {
			"ReadFile": reflect.ValueOf(os.ReadFile),
			"Args":     reflect.ValueOf(os.Args),
			// Add other functions if needed
		},
	})

	// The code to be interpreted
	src := `
        package main

        import (
            "os"
            yaml "gopkg.in/yaml.v3"
        )

        func main() {
            data, err := os.ReadFile("junk.yaml")
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
                println(k, ":", v)
            }
        }
    `

	// Evaluate the code
	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	// Call the main function in the interpreted code
	_, err = i.Eval("main()")
	if err != nil {
		panic(err)
	}
}
