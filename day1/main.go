package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"reflect"
)

func listCreator() []int {
	return 
}

func main() {

    content, err := ioutil.ReadFile("input.txt")

	content := content.Split(result,"\n")

     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(reflect.typeof(content)))
}
