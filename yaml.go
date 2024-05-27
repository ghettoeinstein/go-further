package main

import (
	"fmt"

	"github.com/go-yaml/yaml"
)



func main() {
	// Use the yaml package to marshal the following map to YAML.
	output, err := yaml.Marshal(map[string]string{"a": "b", "c": "d"})
	if err != nil {
        panic(err)
    }
	fmt.Println(output[len(output)-2])
}