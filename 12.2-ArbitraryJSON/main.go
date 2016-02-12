// example of using encoding/json package to read arbitrary JSON file
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// empty map of (string, interface{}) to hold data from json
	var rowMap map[string]interface{}

	// read the row bytes from json file
	fileRowByte, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	// un-marshal json file into map if err != nil
	if err := json.Unmarshal(fileRowByte, &rowMap); err != nil {
		panic(err)
	}

	// print the Name part casted to string
	fmt.Println("Person Name: ", rowMap["name"].(string))
	// print the Age part casted to float64
	fmt.Println("Person Age: ", rowMap["age"].(float64))
	// get the hobbies part as []interface{}
	h := rowMap["hobbies"].([]interface{})
	fmt.Print("Person Hobbies: ")
	// cycle through the interface array and cast each element to string
	for _, i := range h {
		fmt.Print(i.(string), " ")
	}
	fmt.Println()

}
