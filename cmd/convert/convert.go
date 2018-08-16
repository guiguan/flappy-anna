package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	err = ioutil.WriteFile("output.go", []byte(fmt.Sprintf("var Resource_name = %#v\n", data)), 0644)
	check(err)
}
