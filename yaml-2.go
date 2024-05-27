package main

import (
	"fmt"
	"math"

	temp "github.com/go-yaml/yaml"
)


func main() {
	var variableInitializer int = 0
	for i := 0; i < 10; i++ {
	    
		fmt.Println (int(math.Pow(float64(variableInitializer), 2)))
		variableInitializer++
	}
	fmt.Println("Hello lets marshal another set of strings")
	output, _ := temp.Marshal(map[string]string{"a": "b", "c": "d"})
	fmt.Println(string(output))
}