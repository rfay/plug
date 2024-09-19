package main

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	// Import the extracted symbols package
)

func main() {
	i := interp.New(interp.Options{})

	// Use the standard library symbols
	i.Use(stdlib.Symbols)

	// Use the extracted yaml symbols
	i.Use(yaml_yaegi.Symbols)

	// The code to be interpreted
	src := `
        package main

        import (
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
