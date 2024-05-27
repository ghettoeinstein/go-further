package main

import "fmt"

type TestStruct struct {
	Sample string
	SampleInt int
}


func main(){
	 ts := TestStruct{Sample: "Caleb Saunders", SampleInt: 34}

	 fmt.Println(ts.SampleInt)

}

